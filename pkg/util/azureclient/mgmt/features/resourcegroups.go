package features

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"time"

	mgmtfeatures "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-07-01/features"
	"github.com/Azure/go-autorest/autorest"
)

// ResourceGroupsClient is a minimal interface for azure ResourceGroupsClient
type ResourceGroupsClient interface {
	Get(ctx context.Context, resourceGroupName string) (result mgmtfeatures.ResourceGroup, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, parameters mgmtfeatures.ResourceGroup) (result mgmtfeatures.ResourceGroup, err error)
	ResourceGroupsClientAddons
}

type resourceGroupsClient struct {
	mgmtfeatures.ResourceGroupsClient
}

var _ ResourceGroupsClient = &resourceGroupsClient{}

// NewResourceGroupsClient creates a new ResourceGroupsClient
func NewResourceGroupsClient(subscriptionID string, authorizer autorest.Authorizer) ResourceGroupsClient {
	client := mgmtfeatures.NewResourceGroupsClient(subscriptionID)
	client.Authorizer = authorizer
	client.PollingDuration = time.Hour

	return &resourceGroupsClient{
		ResourceGroupsClient: client,
	}
}
