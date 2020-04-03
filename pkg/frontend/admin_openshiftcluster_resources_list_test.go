package frontend

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	mgmtresources "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-07-01/features"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/database"
	"github.com/Azure/ARO-RP/pkg/env"
	"github.com/Azure/ARO-RP/pkg/frontend/kubeactions"
	"github.com/Azure/ARO-RP/pkg/metrics/noop"
	"github.com/Azure/ARO-RP/pkg/util/bucket"
	"github.com/Azure/ARO-RP/pkg/util/clientauthorizer"
	"github.com/Azure/ARO-RP/pkg/util/clusterdata"
	mock_resources "github.com/Azure/ARO-RP/pkg/util/mocks/azureclient/mgmt/resources"
	mock_database "github.com/Azure/ARO-RP/pkg/util/mocks/database"
	utiltls "github.com/Azure/ARO-RP/pkg/util/tls"
	"github.com/Azure/ARO-RP/test/util/listener"
)

func newTestFrontend(ctx context.Context, _env env.Interface, apis map[string]*api.Version, db *database.Database, ka kubeactions.Interface, mr *mock_resources.MockResourcesClient) (Runnable, error) {
	f := &frontend{
		baseLog:         logrus.NewEntry(logrus.StandardLogger()),
		env:             _env,
		db:              db,
		apis:            apis,
		m:               &noop.Noop{},
		kubeActions:     ka,
		resources:       mr,
		bucketAllocator: &bucket.Random{},
	}
	f.ocEnricher = clusterdata.NewBestEffortEnricher(f.baseLog, _env)
	err := f.setupListener(ctx)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func TestAdminListResourcesList(t *testing.T) {
	mockSubID := "00000000-0000-0000-0000-000000000000"
	ctx := context.Background()

	clientkey, clientcerts, err := utiltls.GenerateKeyAndCertificate("client", nil, nil, false, true)
	if err != nil {
		t.Fatal(err)
	}

	serverkey, servercerts, err := utiltls.GenerateKeyAndCertificate("server", nil, nil, false, false)
	if err != nil {
		t.Fatal(err)
	}

	pool := x509.NewCertPool()
	pool.AddCert(servercerts[0])

	cli := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: pool,
				Certificates: []tls.Certificate{
					{
						Certificate: [][]byte{clientcerts[0].Raw},
						PrivateKey:  clientkey,
					},
				},
			},
		},
	}

	type test struct {
		name           string
		resourceID     string
		mocks          func(*test, *mock_database.MockOpenShiftClusters, *mock_resources.MockResourcesClient)
		wantStatusCode int
		wantResponse   func() []byte
		wantError      string
	}

	for _, tt := range []*test{
		{
			name:       "basic coverage",
			resourceID: fmt.Sprintf("/subscriptions/%s/resourcegroups/resourceGroup/providers/Microsoft.RedHatOpenShift/openShiftClusters/resourceName", mockSubID),
			mocks: func(tt *test, openshiftClusters *mock_database.MockOpenShiftClusters, resources *mock_resources.MockResourcesClient) {
				clusterDoc := &api.OpenShiftClusterDocument{
					OpenShiftCluster: &api.OpenShiftCluster{
						ID:   "fakeClusterID",
						Name: "resourceName",
						Type: "Microsoft.RedHatOpenShift/openshiftClusters",
						Properties: api.OpenShiftClusterProperties{
							ClusterProfile: api.ClusterProfile{
								ResourceGroupID: fmt.Sprintf("/subscriptions/%s/resourceGroups/test-cluster", mockSubID),
							},
						},
					},
				}

				openshiftClusters.EXPECT().Get(gomock.Any(), strings.ToLower(tt.resourceID)).
					Return(clusterDoc, nil)

				res := &mgmtresources.ResourceListResult{
					Value: &[]mgmtresources.GenericResourceExpanded{
						{
							Location: to.StringPtr("eastus2"),
							Kind:     to.StringPtr("test2"),
							Tags:     map[string]*string{},
						},
					},
				}
				resources.EXPECT().ListWithDetails(gomock.Any(), "resourceGroup eq 'test-cluster'", "", nil).Return(*res, nil) // TODO
			},
			wantStatusCode: http.StatusOK,
			wantResponse: func() []byte {
				return []byte("{\"value\":[{\"kind\":\"test2\",\"location\":\"eastus2\",\"tags\":{}}]}\n")
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			defer cli.CloseIdleConnections()
			l := listener.NewListener()
			defer l.Close()
			env := &env.Test{
				L:            l,
				TestLocation: "eastus",
				TLSKey:       serverkey,
				TLSCerts:     servercerts,
			}
			env.SetAdminClientAuthorizer(clientauthorizer.NewOne(clientcerts[0].Raw))
			cli.Transport.(*http.Transport).Dial = l.Dial

			controller := gomock.NewController(t)
			defer controller.Finish()

			resourcesClient := mock_resources.NewMockResourcesClient(controller)
			openshiftClusters := mock_database.NewMockOpenShiftClusters(controller)
			tt.mocks(tt, openshiftClusters, resourcesClient)

			f, err := newTestFrontend(ctx, env, api.APIs, &database.Database{
				OpenShiftClusters: openshiftClusters,
			}, nil, resourcesClient)

			if err != nil {
				t.Fatal(err)
			}

			go f.Run(ctx, nil, nil)
			url := fmt.Sprintf("https://server/admin/%s/resources", tt.resourceID)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				t.Fatal(err)
			}
			resp, err := cli.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.wantStatusCode {
				t.Error(resp.StatusCode)
			}

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}

			if tt.wantError == "" {
				if tt.wantResponse != nil {
					if !bytes.Equal(b, tt.wantResponse()) {
						t.Error(string(b))
					}
				}
			} else {
				cloudErr := &api.CloudError{StatusCode: resp.StatusCode}
				err = json.Unmarshal(b, &cloudErr)
				if err != nil {
					t.Fatal(err)
				}

				if cloudErr.Error() != tt.wantError {
					t.Error(cloudErr)
				}
			}
		})
	}
}
