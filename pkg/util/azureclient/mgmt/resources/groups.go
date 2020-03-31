package resources

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"time"

	mgmtresources "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-07-01/features"
	"github.com/Azure/go-autorest/autorest"
)

// GroupsClient is a minimal interface for azure GroupsClient
type GroupsClient interface {
	Get(ctx context.Context, resourceGroupName string) (result mgmtresources.ResourceGroup, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, parameters mgmtresources.ResourceGroup) (result mgmtresources.ResourceGroup, err error)
	GroupsClientAddons
}

type groupsClient struct {
	mgmtresources.ResourceGroupsClient
}

var _ GroupsClient = &groupsClient{}

// NewGroupsClient creates a new ResourcesClient
func NewGroupsClient(subscriptionID string, authorizer autorest.Authorizer) GroupsClient {
	client := mgmtresources.NewResourceGroupsClient(subscriptionID)
	client.Authorizer = authorizer
	client.PollingDuration = time.Hour

	return &groupsClient{
		ResourceGroupsClient: client,
	}
}
