package resources

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	mgmtresources "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-07-01/features"
)

// ResourcesClientAddons is a minimal interface for azure ResourcesClient
type ResourcesClientAddons interface {
	ListWithDetails(ctx context.Context, filter, expand string, top *int32) (mgmtresources.ResourceListResult, error)
}

func (c *resourcesClient) ListWithDetails(ctx context.Context, filter, expand string, top *int32) (mgmtresources.ResourceListResult, error) {
	page, err := c.ResourcesClient.List(ctx, filter, expand, top)
	result := mgmtresources.ResourceListResult{Value: &[]mgmtresources.GenericResourceExpanded{}}
	if err != nil {
		return result, err
	}
	for page.NotDone() {
		for _, res := range page.Values() {
			if res.ID != nil {
				gr, err := c.ResourcesClient.GetByID(ctx, *res.ID, "2019-07-01")
				if err == nil {
					res.Name = gr.Name
					res.Type = gr.Type
					res.Location = gr.Location
					res.Kind = gr.Kind
					res.Identity = gr.Identity
					res.Properties = gr.Properties
					res.Plan = gr.Plan
					res.Tags = gr.Tags
					res.Sku = gr.Sku
					res.ManagedBy = gr.ManagedBy
				}
			}
			*result.Value = append(*result.Value, res)
		}

		err = page.NextWithContext(ctx)
		if err != nil {
			return result, err
		}
	}

	return result, nil
}
