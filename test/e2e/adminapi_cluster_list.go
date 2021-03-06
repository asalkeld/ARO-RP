package e2e

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Azure/ARO-RP/pkg/api/admin"
)

var _ = Describe("[Admin API] List clusters action", func() {
	BeforeEach(skipIfNotInDevelopmentEnv)

	It("should be able to return list of all clusters with admin fields", func() {
		ctx := context.Background()
		resourceID := resourceIDFromEnv()

		// TODO: add pagination support once https://github.com/Azure/ARO-RP/pull/859 lands.
		//       Replace code below with testAdminClustersList(...). See example below.
		By("requesting the cluster document via RP admin API")
		var ocs []*admin.OpenShiftCluster
		resp, err := adminRequest(ctx, http.MethodGet, "/admin/providers/Microsoft.RedHatOpenShift/openShiftClusters", nil, nil, &ocs)
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))

		By("checking that we received the expected cluster")
		var oc *admin.OpenShiftCluster
		for i := range ocs {
			if ocs[i].ID == resourceID {
				oc = ocs[i]
			}
		}
		Expect(oc).ToNot(BeNil())

		By("checking that fields available only in Admin API have values")
		// Note: some fields will have empty values
		// on successfully provisioned cluster (oc.Properties.Install, for example)
		Expect(oc.Properties.ServicePrincipalProfile.TenantID).ToNot(BeEmpty())
		Expect(oc.Properties.StorageSuffix).ToNot(BeEmpty())
	})

	It("should be able to return list clusters with admin fields by subscription", func() {
		ctx := context.Background()
		resourceID := resourceIDFromEnv()
		subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")

		path := fmt.Sprintf("/subscriptions/%s/providers/Microsoft.RedHatOpenShift/openShiftClusters", subscriptionID)
		testAdminClustersList(ctx, path, resourceID)
	})

	It("should be able to return list clusters with admin fields by resource group", func() {
		ctx := context.Background()
		resourceID := resourceIDFromEnv()
		subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
		resourceGroup := os.Getenv("RESOURCEGROUP")

		path := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RedHatOpenShift/openShiftClusters", subscriptionID, resourceGroup)
		testAdminClustersList(ctx, path, resourceID)
	})
})

func testAdminClustersList(ctx context.Context, path, wantResourceID string) {
	By("requesting the cluster document via RP admin API")
	ocs := make([]*admin.OpenShiftCluster, 0)
	params := url.Values{}
	for {
		var list admin.OpenShiftClusterList
		resp, err := adminRequest(ctx, http.MethodGet, path, params, nil, &list)
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))

		ocs = append(ocs, list.OpenShiftClusters...)

		if list.NextLink == "" {
			break
		}

		params = nextParams(list.NextLink)
	}

	By("checking that we received the expected cluster")
	var oc *admin.OpenShiftCluster
	for i := range ocs {
		if ocs[i].ID == wantResourceID {
			oc = ocs[i]
		}
	}
	Expect(oc).ToNot(BeNil())

	By("checking that fields available only in Admin API have values")
	// Note: some fields will have empty values
	// on successfully provisioned cluster (oc.Properties.Install, for example)
	Expect(oc.Properties.ServicePrincipalProfile.TenantID).ToNot(BeEmpty())
	Expect(oc.Properties.StorageSuffix).ToNot(BeEmpty())
}

func nextParams(nextLink string) url.Values {
	url, err := url.Parse(nextLink)
	Expect(err).NotTo(HaveOccurred())

	return url.Query()
}
