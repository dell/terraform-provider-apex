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
	_ datasource.DataSource              = &clonesDataSource{}
	_ datasource.DataSourceWithConfigure = &clonesDataSource{}
)

// NewClonesDataSource returns a new clones data source instance.
func NewClonesDataSource() datasource.DataSource {
	return &clonesDataSource{}
}

// clonesDataSource is the data source implementation.
type clonesDataSource struct {
	client *client.APIClient
}

// Metadata returns the data source type name.
func (d *clonesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_navigator_block_clones"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (d *clonesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "This Terraform Datasource is used to query existing clones on Apex Navigator." +
			" The information fetched from this block can be further used for resource block.",
		MarkdownDescription: "This Terraform Datasource is used to query existing clones on Apex Navigator." +
			" The information fetched from this block can be further used for resource block.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "ID of the clone datasource",
				MarkdownDescription: "ID of the clone datasource",
				Computed:            true,
			},
			"clones": schema.ListNestedAttribute{
				Description:         "List of clones",
				MarkdownDescription: "List of clones",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: ClonesDataSourceSchema.Attributes,
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
func (d *clonesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.ClonesDataSourceModel
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

	clones, status, err := helper.GetCloneCollection(d.client, filter)
	if (err != nil) || (status.StatusCode != http.StatusOK && status.StatusCode != http.StatusPartialContent) {
		resp.Diagnostics.AddError(
			constants.BlockCloneReadErrorMsg,
			err.Error(),
		)
		return
	}

	// If the returned filtered values does not equal the length of the filter
	// Then one or more of the filtered values are invalid
	if filterUsed && len(clones.Results) != len(state.Filter.IDs) {
		resp.Diagnostics.AddError(
			constants.BlockCloneFilterErrorMsg,
			constants.FilterGeneralErrorMsg,
		)
		return
	}

	// Map response body to model
	for _, clone := range clones.Results {
		cloneState := helper.GetClonesModelDs(clone)
		state.Clones = append(state.Clones, cloneState)
	}
	state.ID = types.StringValue("clone-ds-id")

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *clonesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
