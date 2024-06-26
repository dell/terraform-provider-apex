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
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &mobilityGroupsDataSource{}
	_ datasource.DataSourceWithConfigure = &mobilityGroupsDataSource{}
)

// NewMobilityGroupsDataSource is a storage system data source object
func NewMobilityGroupsDataSource() datasource.DataSource {
	return &mobilityGroupsDataSource{}
}

// mobilityGroupsDataSource is the data source implementation.
type mobilityGroupsDataSource struct {
	client *client.APIClient
}

func (d *mobilityGroupsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_navigator_block_mobility_groups"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (d *mobilityGroupsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) { // nolint:funlen, dupl
	resp.Schema = schema.Schema{
		Description: "This Terraform Datasource is used to query existing mobility groups(source) on Apex Navigator." +
			" The information fetched from this block can be further used for resource block.",
		MarkdownDescription: "This Terraform Datasource is used to query existing mobility groups(source) on Apex Navigator." +
			" The information fetched from this block can be further used for resource block.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "ID of the Mobility Group datasource",
				MarkdownDescription: "ID of the Mobility Group datasource",
				Computed:            true,
			},
			"mobility_groups": schema.ListNestedAttribute{
				Description:         "List of mobility groups(source) available on Apex Navigator.",
				MarkdownDescription: "List of mobility groups(source) available on Apex Navigator.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: MobilityGroupsDataSourceSchema.Attributes,
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
func (d *mobilityGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) { // nolint:gocognit, funlen
	var state models.MobilityGroupsDataSourceModel
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

	mobilityGroups, status, err := helper.GetMobilityGroupCollection(d.client, filter)
	if (err != nil) || (status.StatusCode != http.StatusOK && status.StatusCode != http.StatusPartialContent) {
		resp.Diagnostics.AddError(
			constants.BlockMobilityGroupReadErrorMsg,
			err.Error(),
		)
		return
	}

	// If the returned filtered values does not equal the length of the filter
	// Then one or more of the filtered values are invalid
	if filterUsed && len(mobilityGroups.Results) != len(state.Filter.IDs) {
		resp.Diagnostics.AddError(
			constants.BlockMobilityGroupFilterErrorMsg,
			constants.FilterGeneralErrorMsg,
		)
		return
	}

	// Map response body to model
	for _, mobilityGroup := range mobilityGroups.Results {
		mobilityGroupsState := helper.GetMobilityGroupModelDs(mobilityGroup)

		state.MobilityGroups = append(state.MobilityGroups, mobilityGroupsState)
	}
	state.ID = types.StringValue("mobility-group-ds-id")

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *mobilityGroupsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
