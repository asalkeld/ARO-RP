package frontend

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"regexp"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/database"
	"github.com/Azure/ARO-RP/pkg/env"
	"github.com/Azure/ARO-RP/pkg/frontend/kubeactions"
	"github.com/Azure/ARO-RP/pkg/metrics/noop"
	"github.com/Azure/ARO-RP/pkg/util/bucket"
	"github.com/Azure/ARO-RP/pkg/util/clusterdata"
	mockfeatures "github.com/Azure/ARO-RP/pkg/util/mocks/azureclient/mgmt/features"
)

func newTestFrontend(ctx context.Context, _env env.Interface, apis map[string]*api.Version, db *database.Database, ka kubeactions.Interface, mr *mockfeatures.MockResourcesClient) (Runnable, error) {
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

func TestURLPathsAreLowerCase(t *testing.T) {
	f := &frontend{
		baseLog: logrus.NewEntry(logrus.StandardLogger()),
	}
	router := f.setupRouter()

	varCleanupRe := regexp.MustCompile(`{.*?}`)
	err := router.Walk(func(route *mux.Route, _ *mux.Router, _ []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err != nil {
			// Ignore the error: it can occur when a route has no path,
			// but there is no way to check it here
			return nil
		}

		cleanPathTemplate := varCleanupRe.ReplaceAllString(pathTemplate, "")
		if cleanPathTemplate != strings.ToLower(cleanPathTemplate) {
			t.Error(pathTemplate)
		}

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
