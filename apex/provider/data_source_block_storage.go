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

	client "github.com/dell/terraform-provider-apex/client/apexclient/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/dell/terraform-provider-apex/apex/models"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &blockStoragesDataSource{}
	_ datasource.DataSourceWithConfigure = &blockStoragesDataSource{}
)

// NewBlockStoragesDataSource is a block storage data source object
func NewBlockStoragesDataSource() datasource.DataSource {
	return &blockStoragesDataSource{}
}

// blockStorageDataSource is the data source implementation.
type blockStoragesDataSource struct {
	client *client.APIClient
}

func (d *blockStoragesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_block_storages"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (d *blockStoragesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) { // nolint:funlen
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"block_storages": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: BlockStorageDataSourceSchema.Attributes,
				},
			},
		},
		Blocks: map[string]schema.Block{
			"filter": schema.SingleNestedBlock{
				Attributes: map[string]schema.Attribute{
					"ids": schema.SetAttribute{
						Description:         "Filter by ids",
						MarkdownDescription: "Filter by ids",
						Optional:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}

// Read method is used to refresh the Terraform state based on the schema data.
func (d *blockStoragesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) { // nolint:gocognit, funlen
	var state models.BlockStoragesDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check that the filter is valid
	filter := ""
	filterUsed := false
	if state.Filter != nil && state.Filter.IDs != nil && len(state.Filter.IDs) > 0 {
		filterUsed = true
		filteredNames := make([]string, 0)
		for _, ids := range state.Filter.IDs {
			filteredNames = append(filteredNames, ids.ValueString())
		}
		filter = helper.CreateFilter(filteredNames, "id")
	}

	storageSystems, status, err := helper.GetBlockStorageCollection(d.client, filter)
	if (err != nil) || (status.StatusCode != http.StatusOK && status.StatusCode != http.StatusPartialContent) {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Unable to Read Apex Navigator Block Storage: "+newErr,
			err.Error(),
		)
		return
	}

	// If the returned filtered values does not equal the length of the filter
	// Then one or more of the filtered values are invalid
	if filterUsed && len(storageSystems.Results) != len(state.Filter.IDs) {
		resp.Diagnostics.AddError(
			"Failed to filter block storage.",
			"one more more of the ids set in the filter is invalid.",
		)
		return
	}

	// needs nested for loop for deployment details
	// Map response body to model
	for _, storageSystem := range storageSystems.Results {
		storageSystemsState := helper.GetBlockStorageSystem(storageSystem)
		if storageSystem.DeploymentDetails != nil {
			if storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails != nil {
				storageSystemsState.DeploymentDetails.SystemPublicCloud.VirtualPrivateCloud = types.StringPointerValue(storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails.VirtualPrivateCloud)
			}
		} else {
			resp.Diagnostics.AddError(
				"Unexpected system type",
				"Unexpected system type",
			)
		}

		state.BlockStorages = append(state.BlockStorages, storageSystemsState)
	}
	state.ID = types.StringValue("block-storage-ds-id")

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *blockStoragesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
