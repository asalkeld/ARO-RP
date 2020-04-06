package features

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	mgmtfeatures "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-07-01/features"
)

// ResourcesClientAddons is a minimal interface for azure ResourcesClient
type ResourcesClientAddons interface {
	List(ctx context.Context, filter, expand string, top *int32) (mgmtfeatures.ResourceListResult, error)
}

func (c *resourcesClient) List(ctx context.Context, filter, expand string, top *int32) (mgmtfeatures.ResourceListResult, error) {
	page, err := c.ResourcesClient.List(ctx, filter, expand, top)
	result := mgmtfeatures.ResourceListResult{Value: &[]mgmtfeatures.GenericResourceExpanded{}}
	if err != nil {
		return result, err
	}
	for page.NotDone() {
		*result.Value = append(*result.Value, page.Values()...)
		err = page.NextWithContext(ctx)
		if err != nil {
			return result, err
		}
	}

	return result, nil
}
