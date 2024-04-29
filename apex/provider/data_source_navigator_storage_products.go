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

package provider

import (
	"context"
	"fmt"
	"net/http"

	client "dell/apex-client"

	"github.com/dell/terraform-provider-apex/apex/constants"
	"github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/dell/terraform-provider-apex/apex/models"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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

func (d *storageProductsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_navigator_storage_products"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (d *storageProductsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This Terraform Datasource is used to query existing storage products on Apex Navigator.",
		MarkdownDescription: "This Terraform Datasource is used to query existing storage products on Apex Navigator.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "ID of the storage products datasource",
				MarkdownDescription: "ID of the storage products datasource",
				Computed:            true,
			},
			"storage_products": schema.ListNestedAttribute{
				Computed:            true,
				Description:         "List of storage products",
				MarkdownDescription: "List of storage products",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description:         "Identifier of the storage product.",
							MarkdownDescription: "Identifier of the storage product.",
							Optional:            true,
						},
						"name": schema.StringAttribute{
							Description:         "Name of the storage product.",
							MarkdownDescription: "Name of the storage product.",
							Optional:            true,
						},
						"system_type": schema.StringAttribute{
							Description:         "Enum for all the supported storage products (e.g., POWERFLEX, APEX block storage for cloud).",
							MarkdownDescription: "Enum for all the supported storage products (e.g., POWERFLEX, APEX block storage for cloud).",
							Optional:            true,
						},
						"storage_type": schema.StringAttribute{
							Description:         "Enum representing all the supported storage types (e.g., BLOCK).",
							MarkdownDescription: "Enum representing all the supported storage types (e.g., BLOCK).",
							Optional:            true,
						},
						"description": schema.StringAttribute{
							Description:         "Description of the storage product and its capabilities.",
							MarkdownDescription: "Description of the storage product and its capabilities.",
							Optional:            true,
						},
						"cloud_type": schema.StringAttribute{
							Description:         "Enum for all the supported cloud providers (e.g., AWS - Amazon Web Services).",
							MarkdownDescription: "Enum for all the supported cloud providers (e.g., AWS - Amazon Web Services).",
							Optional:            true,
						},
						"latest_version": schema.StringAttribute{
							Description:         "Latest supported version of the storage product on the cloud.",
							MarkdownDescription: "Latest supported version of the storage product on the cloud.",
							Optional:            true,
						},

						"support_map": schema.ListNestedAttribute{
							Description:         "Array of support mappings for the storage product.",
							MarkdownDescription: "Array of support mappings for the storage product.",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "Identifier of the support mapping.",
										MarkdownDescription: "Identifier of the support mapping.",
										Optional:            true,
									},
									"supported_evaluation_period": schema.Int64Attribute{
										Description:         "Evaluation period in days. After the evaluation period is expired, you need to purchase a license from Dell, to continue using the product.",
										MarkdownDescription: "Evaluation period in days. After the evaluation period is expired, you need to purchase a license from Dell, to continue using the product.",
										Optional:            true,
									},
									"version": schema.StringAttribute{
										Optional:            true,
										Description:         "Version of the storage product on the cloud.",
										MarkdownDescription: "Version of the storage product on the cloud.",
									},
									"supported_actions": schema.ListNestedAttribute{
										Description:         "Enum for all the supported storage products actions (e.g., CREATE, DELETE, DEPLOY).",
										MarkdownDescription: "Enum for all the supported storage products actions (e.g., CREATE, DELETE, DEPLOY).",
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
		},
		Blocks: map[string]schema.Block{
			"filter": schema.SingleNestedBlock{
				Attributes: map[string]schema.Attribute{
					"system_type": schema.StringAttribute{
						Description:         "Filter by system_type",
						MarkdownDescription: "Filter by system_type",
						Optional:            true,
					},
				},
			},
		},
	}
}

// Read method is used to refresh the Terraform state based on the schema data.
func (d *storageProductsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.StorageProductsDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check that the filter is valid
	filter := ""
	filterUsed := false
	if state.Filter != nil && state.Filter.SystemType.ValueString() != "" {
		filterUsed = true
		filteredNames := make([]string, 0)
		filteredNames = append(filteredNames, state.Filter.SystemType.ValueString())
		filter = helper.CreateFilter(filteredNames, "system_type")
	}

	storageProducts, status, err := helper.GetStorageProductsCollection(d.client, filter)
	if (err != nil) || (status.StatusCode != http.StatusOK && status.StatusCode != http.StatusPartialContent) {
		errMsg := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.StorageProductReadErrorMsg,
			errMsg,
		)
		return
	}

	// If the returned filtered values does not equal the length of the filter
	// Then one or more of the filtered values are invalid
	if filterUsed && len(storageProducts.Results) != 1 {
		tflog.Info(ctx, fmt.Sprintf("Filtered %v storage products", storageProducts.Results))
		resp.Diagnostics.AddError(
			constants.FilterErrorSystemTypeMsg,
			constants.StorageProductFilterErrorMsg,
		)
		return
	}

	state = helper.MapStorageProduct(storageProducts)

	state.ID = types.StringValue("storage-product-ds-id")

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
			constants.DatasourceConfigureErrorMsg,
			fmt.Sprintf(constants.GeneralConfigureErrorMsg, req.ProviderData),
		)

		return
	}

	d.client = client
}
