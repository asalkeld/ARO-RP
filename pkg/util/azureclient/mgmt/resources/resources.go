package resources

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"time"

	mgmtresources "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-07-01/features"
	"github.com/Azure/go-autorest/autorest"
)

// ResourcesClient is a minimal interface for azure ResourcesClient
type ResourcesClient interface {
	ResourcesClientAddons
}

type resourcesClient struct {
	mgmtresources.ResourcesClient
}

var _ ResourcesClient = &resourcesClient{}

// NewResourcesClient creates a new ResourcesClient
func NewResourcesClient(subscriptionID string, authorizer autorest.Authorizer) ResourcesClient {
	client := mgmtresources.NewResourcesClient(subscriptionID)
	client.Authorizer = authorizer
	client.PollingDuration = time.Hour
	client.PollingDelay = 10 * time.Second

	return &resourcesClient{
		ResourcesClient: client,
	}
}
