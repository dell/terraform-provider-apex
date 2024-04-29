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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/dell/terraform-provider-apex/apex/constants"
	"github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/dell/terraform-provider-apex/apex/models"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &storagesDataSource{}
	_ datasource.DataSourceWithConfigure = &storagesDataSource{}
)

// NewStoragesDataSource is a storage data source object
func NewStoragesDataSource() datasource.DataSource {
	return &storagesDataSource{}
}

// StorageDataSource is the data source implementation.
type storagesDataSource struct {
	client *client.APIClient
}

func (d *storagesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_navigator_storages"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (d *storagesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) { // nolint:funlen
	resp.Schema = schema.Schema{
		Description:         "This Terraform Datasource is used to query existing storages on Apex Navigator.",
		MarkdownDescription: "This Terraform Datasource is used to query existing storages on Apex Navigator.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "ID of the storage datasource",
				MarkdownDescription: "ID of the storage datasource",
				Computed:            true,
			},
			"storages": schema.ListNestedAttribute{
				Description:         "List of storages",
				MarkdownDescription: "List of storages",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: StorageDataSourceSchema.Attributes,
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
					"system_type": schema.StringAttribute{
						Description:         "Filter by system type",
						MarkdownDescription: "Filter by system type",
						Optional:            true,
						Validators: []validator.String{
							stringvalidator.OneOf("ISILON", "POWERFLEX"),
						},
					},
				},
			},
		},
	}
}

// Read method is used to refresh the Terraform state based on the schema data.
func (d *storagesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) { // nolint:gocognit, funlen
	var state models.StoragesDataSourceModel
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

	if state.Filter != nil && state.Filter.SystemType.ValueString() != "" {
		filterUsed = true
		systemFilter := helper.CreateFilter([]string{state.Filter.SystemType.ValueString()}, "system_type")
		// If there is a ID filter and a system type filter append the system type filter to the id filter
		if filter != "" {
			filter = fmt.Sprintf("%s or %s", filter, systemFilter)
		} else { // Else just set the system type filter
			filter = systemFilter
		}

	}

	storageSystems, status, err := helper.GetStorageCollection(d.client, filter)
	if (err != nil) || (status.StatusCode != http.StatusOK && status.StatusCode != http.StatusPartialContent) {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.StorageReadErrorMsg,
			newErr,
		)
		return
	}

	// If the returned filtered values does not equal the length of the filter
	// Then one or more of the filtered values are invalid
	if filterUsed && state.Filter.SystemType.ValueString() == "" && len(storageSystems.Results) != len(state.Filter.IDs) {
		resp.Diagnostics.AddError(
			constants.StorageFilterErrorMsg,
			constants.FilterGeneralErrorMsg,
		)
		return
	}

	// needs nested for loop for deployment details
	// Map response body to model
	for _, storageSystem := range storageSystems.Results {
		storageSystemsState := helper.GetStorageSystem(storageSystem)
		if storageSystem.DeploymentDetails != nil {
			if storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails != nil {
				storageSystemsState.DeploymentDetails.SystemPublicCloud.VirtualPrivateCloud = types.StringPointerValue(storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails.VirtualPrivateCloud)
			}
		} else {
			resp.Diagnostics.AddError(
				constants.UnexpectedSysteType,
				constants.UnexpectedSysteType,
			)
		}

		state.Storages = append(state.Storages, storageSystemsState)
	}
	state.ID = types.StringValue("storage-ds-id")

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *storagesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
