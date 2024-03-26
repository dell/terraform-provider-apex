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

	"github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/dell/terraform-provider-apex/apex/models"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &volumesDataSource{}
	_ datasource.DataSourceWithConfigure = &volumesDataSource{}
)

// NewVolumesDataSource returns a new volumes data source instance.
func NewVolumesDataSource() datasource.DataSource {
	return &volumesDataSource{}
}

// volumesDataSource is the data source implementation.
type volumesDataSource struct {
	client *client.APIClient
}

// Metadata returns the data source type name.
func (d *volumesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_volumes"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (d *volumesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"volumes": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"system_id": schema.StringAttribute{
							Computed: true,
						},
						"system_type": schema.StringAttribute{
							Computed: true,
						},
						"allocated_size": schema.Int64Attribute{
							Computed: true,
						},
						"bandwidth": schema.Int64Attribute{
							Computed: true,
						},
						"consistency_group_name": schema.StringAttribute{
							Computed: true,
						},
						"data_reduction_percent": schema.Float64Attribute{
							Computed: true,
						},
						"data_reduction_ratio": schema.Float64Attribute{
							Computed: true,
						},
						"data_reduction_saved_size": schema.Int64Attribute{
							Computed: true,
						},
						"io_limit_policy_name": schema.StringAttribute{
							Computed: true,
						},
						"iops": schema.Int64Attribute{
							Computed: true,
						},
						"is_compressed_or_deduped": schema.StringAttribute{
							Computed: true,
						},
						"is_thin_enabled": schema.BoolAttribute{
							Computed: true,
						},
						"issue_count": schema.Int64Attribute{
							Computed: true,
						},
						"latency": schema.Int64Attribute{
							Computed: true,
						},
						"logical_size": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"native_id": schema.StringAttribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Computed: true,
						},
						"pool_id": schema.StringAttribute{
							Computed: true,
						},
						"pool_name": schema.StringAttribute{
							Computed: true,
						},
						"pool_type": schema.StringAttribute{
							Computed: true,
						},
						"snap_shot_count": schema.Int64Attribute{
							Computed: true,
						},
						"snap_shot_policy": schema.StringAttribute{
							Computed: true,
						},
						"snap_shot_size": schema.Int64Attribute{
							Computed: true,
						},
						"storage_resource_id": schema.StringAttribute{
							Computed: true,
						},
						"storage_resource_native_id": schema.StringAttribute{
							Computed: true,
						},
						"system_model": schema.StringAttribute{
							Computed: true,
						},
						"system_name": schema.StringAttribute{
							Computed: true,
						},
						"total_size": schema.Int64Attribute{
							Computed: true,
						},
						"used_size": schema.Int64Attribute{
							Computed: true,
						},
						"used_size_unique": schema.Int64Attribute{
							Computed: true,
						},
					},
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
func (d *volumesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.VolumesDataSourceModel
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

	volumes, status, err := helper.GetVolumesCollection(d.client, filter)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Apex Navigator Volumes",
			err.Error(),
		)
		return
	}
	if status.StatusCode != http.StatusOK && status.StatusCode != http.StatusPartialContent {
		resp.Diagnostics.AddError(
			"Unable to Read Apex Navigator Volumes",
			status.Status,
		)
		return
	}

	// If the returned filtered values does not equal the length of the filter
	// Then one or more of the filtered values are invalid
	if filterUsed && len(volumes.Results) != len(state.Filter.IDs) {
		resp.Diagnostics.AddError(
			"Failed to filter mobility groups.",
			"one more more of the ids set in the filter is invalid.",
		)
		return
	}

	// Map response body to model
	state.Volumes = helper.MapVolumesToState(volumes.GetResults())
	state.ID = types.StringValue("volumes-id")

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *volumesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *openapi.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}
