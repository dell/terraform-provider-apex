/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package provider handles resource and data sources.
package provider

import (
	"context"
	"fmt"
	"net/http"

	client "github.com/dell/terraform-provider-apex/pkg/gen/apex/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &storageProductsDataSource{}
	_ datasource.DataSourceWithConfigure = &storageProductsDataSource{}
)

// NewStorageProductsDataSource returns a new storage products data source instance.
func NewStorageProductsDataSource() datasource.DataSource {
	return &storageProductsDataSource{}
}

// storageProductsDataSource is the data source implementation.
type storageProductsDataSource struct {
	client *client.APIClient
}

// storageProductsDataSourceModel maps the data source schema data.
type storageProductsDataSourceModel struct {
	StorageProducts []storageProductsModel `tfsdk:"storage_products"`
	ID              types.String           `tfsdk:"id"`
}

// storageProductsModel maps hosts schema data.
type storageProductsModel struct {
	ID            types.String                 `tfsdk:"id"`
	Name          types.String                 `tfsdk:"name"`
	SystemType    client.StorageSystemTypeEnum `tfsdk:"system_type"`
	StorageType   client.StorageTypeEnum       `tfsdk:"storage_type"`
	Description   types.String                 `tfsdk:"description"`
	CloudType     client.CloudProviderEnum     `tfsdk:"cloud_type"`
	LatestVersion types.String                 `tfsdk:"latest_version"`
	SupportMaps   []supportMapModel            `tfsdk:"support_map"`
}

// supportMapModel maps supportMap schema data
type supportMapModel struct {
	ID                        types.String                      `tfsdk:"id"`
	SupportedEvaluationPeriod types.Int64                       `tfsdk:"supported_evaluation_period"`
	Version                   types.String                      `tfsdk:"version"`
	SupportedActions          []client.StorageProductActionEnum `tfsdk:"supported_actions"`
}

func (d *storageProductsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_storage_products"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (d *storageProductsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"storage_products": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "",
							Optional:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the storage proct",
							Optional:            true,
						},
						"system_type": schema.StringAttribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"storage_type": schema.StringAttribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Description of the storage product and its capabilities",
							Optional:            true,
						},
						"cloud_type": schema.StringAttribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"latest_version": schema.StringAttribute{
							MarkdownDescription: "Latest supported version of the storage product on the cloud",
							Optional:            true,
						},

						"support_map": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Optional: true,
									},
									"supported_evaluation_period": schema.Int64Attribute{
										Optional: true,
									},
									"version": schema.StringAttribute{
										Optional: true,
									},
									"supported_actions": schema.ListNestedAttribute{
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// Read method is used to refresh the Terraform state based on the schema data.
func (d *storageProductsDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state storageProductsDataSourceModel

	storageProducts, status, err := d.client.StorageProductsAPI.StorageProductsCollection(context.Background()).Limit(500).Execute()
	if (err != nil) || (status.StatusCode != http.StatusOK && status.StatusCode != http.StatusPartialContent) {
		resp.Diagnostics.AddError(
			"Unable Read to Storage Product Volume",
			err.Error(),
		)
		return
	}
	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}

	// Map response body to model
	for _, storageProduct := range storageProducts.Results {
		storageProductState := storageProductsModel{
			ID:            types.StringValue(storageProduct.Id),
			Name:          types.StringValue(storageProduct.Name),
			Description:   types.StringValue(storageProduct.Description),
			LatestVersion: types.StringValue(storageProduct.LatestVersion),
			SystemType:    client.StorageSystemTypeEnum(storageProduct.SystemType),
			StorageType:   storageProduct.StorageType,
			CloudType:     storageProduct.CloudType,
		}

		for _, supportmap := range storageProduct.SupportMap {
			storageProductState.SupportMaps = append(storageProductState.SupportMaps, supportMapModel{
				ID:                        types.StringValue(supportmap.Id),
				SupportedEvaluationPeriod: types.Int64Value(int64(supportmap.SupportedEvaluationPeriod)),
				Version:                   types.StringValue(supportmap.Version),
				SupportedActions:          []client.StorageProductActionEnum{},
			})
		}

		state.StorageProducts = append(state.StorageProducts, storageProductState)
	}

	state.ID = types.StringValue("placeholder")

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *storageProductsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *openapi.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}
