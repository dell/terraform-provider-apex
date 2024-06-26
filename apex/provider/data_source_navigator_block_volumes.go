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
	resp.TypeName = req.ProviderTypeName + "_navigator_block_volumes"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (d *volumesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This Terraform Datasource is used to query existing volumes on Apex Navigator.",
		MarkdownDescription: "This Terraform Datasource is used to query existing volumes on Apex Navigator.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "ID of the volume datasource",
				MarkdownDescription: "ID of the volume datasource",
				Computed:            true,
			},
			"volumes": schema.ListNestedAttribute{
				Computed:            true,
				Description:         "List of volumes",
				MarkdownDescription: "List of volumes",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description:         "Identifier of the volume",
							MarkdownDescription: "Identifier of the volume",
							Computed:            true,
						},
						"system_id": schema.StringAttribute{
							Description:         "Unique identifier for the device or appliance.",
							MarkdownDescription: "Unique identifier for the device or appliance.",
							Computed:            true,
						},
						"system_type": schema.StringAttribute{
							Description:         "Type of the system for the volume",
							MarkdownDescription: "Type of the system for the volume",
							Computed:            true,
						},
						"allocated_size": schema.Int64Attribute{
							Description:         "The allocated size of the volume - Unit: bytes",
							MarkdownDescription: "The allocated size of the volume - Unit: bytes",
							Computed:            true,
						},
						"bandwidth": schema.Int64Attribute{
							Description:         "The bandwidth consumed by the volume. Aggregated for a rolling average over the last 24 hours - Unit: bytes/s",
							MarkdownDescription: "The bandwidth consumed by the volume. Aggregated for a rolling average over the last 24 hours - Unit: bytes/s",
							Computed:            true,
						},
						"consistency_group_name": schema.StringAttribute{
							Description:         "Consistency group name of the volume.",
							MarkdownDescription: "Consistency group name of the volume.",
							Computed:            true,
						},
						"data_reduction_percent": schema.Float64Attribute{
							Description:         "The data reduction percent for the volume.",
							MarkdownDescription: "The data reduction percent for the volume.",
							Computed:            true,
						},
						"data_reduction_ratio": schema.Float64Attribute{
							Description:         "The data reduction ratio for the volume.",
							MarkdownDescription: "The data reduction ratio for the volume.",
							Computed:            true,
						},
						"data_reduction_saved_size": schema.Int64Attribute{
							Description:         "The data reduction capacity saved for the volume - Unit: bytes",
							MarkdownDescription: "The data reduction capacity saved for the volume - Unit: bytes",
							Computed:            true,
						},
						"io_limit_policy_name": schema.StringAttribute{
							Description:         "The IO limit policy name for the volume.",
							MarkdownDescription: "The IO limit policy name for the volume.",
							Computed:            true,
						},
						"iops": schema.Int64Attribute{
							Description:         "The IOPS for the volume. Aggregated for a rolling average over the last 24 hours - Unit: IO/s",
							MarkdownDescription: "The IOPS for the volume. Aggregated for a rolling average over the last 24 hours - Unit: IO/s",
							Computed:            true,
						},
						"is_compressed_or_deduped": schema.StringAttribute{
							Description:         "Identifies whether the volume is compressed or deduplicated.",
							MarkdownDescription: "Identifies whether the volume is compressed or deduplicated.",
							Computed:            true,
						},
						"is_thin_enabled": schema.BoolAttribute{
							Description:         "Identifies whether the volume has thin provisioning enabled.",
							MarkdownDescription: "Identifies whether the volume has thin provisioning enabled.",
							Computed:            true,
						},
						"issue_count": schema.Int64Attribute{
							Description:         "Number of health issues that are present on the volume.",
							MarkdownDescription: "Number of health issues that are present on the volume.",
							Computed:            true,
						},
						"latency": schema.Int64Attribute{
							Description:         "The latency for the volume. Aggregated for a rolling average over the last 24 hours - Unit: microseconds",
							MarkdownDescription: "The latency for the volume. Aggregated for a rolling average over the last 24 hours - Unit: microseconds",
							Computed:            true,
						},
						"logical_size": schema.Int64Attribute{
							Description:         "The logical size for the volume - Unit: bytes",
							MarkdownDescription: "The logical size for the volume - Unit: bytes",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							Description:         "The name of the volume.",
							MarkdownDescription: "The name of the volume.",
							Computed:            true,
						},
						"native_id": schema.StringAttribute{
							Description:         "Identifier of the volume, defined by the system.",
							MarkdownDescription: "Identifier of the volume, defined by the system.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							Description:         "Type of the volume, which is either LUN or VOLUME.",
							MarkdownDescription: "Type of the volume, which is either LUN or VOLUME.",
							Computed:            true,
						},
						"pool_id": schema.StringAttribute{
							Description:         "The pool identifier for the volume.",
							MarkdownDescription: "The pool identifier for the volume.",
							Computed:            true,
						},
						"pool_name": schema.StringAttribute{
							Description:         "The pool name for the volume.",
							MarkdownDescription: "The pool name for the volume.",
							Computed:            true,
						},
						"pool_type": schema.StringAttribute{
							Description:         "Type of the pool.",
							MarkdownDescription: "Type of the pool.",
							Computed:            true,
						},
						"snap_shot_count": schema.Int64Attribute{
							Description:         "The snapshot count for the volume.",
							MarkdownDescription: "The snapshot count for the volume.",
							Computed:            true,
						},
						"snap_shot_policy": schema.StringAttribute{
							Description:         "The snapshot policy for the volume.",
							MarkdownDescription: "The snapshot policy for the volume.",
							Computed:            true,
						},
						"snap_shot_size": schema.Int64Attribute{
							Description:         "The snapshot size for the volume - Unit: bytes",
							MarkdownDescription: "The snapshot size for the volume - Unit: bytes",
							Computed:            true,
						},
						"storage_resource_id": schema.StringAttribute{
							Description:         "The storage resource identifier for the volume.",
							MarkdownDescription: "The storage resource identifier for the volume.",
							Computed:            true,
						},
						"storage_resource_native_id": schema.StringAttribute{
							Description:         "The storage resource native identifier for the volume.",
							MarkdownDescription: "The storage resource native identifier for the volume.",
							Computed:            true,
						},
						"system_model": schema.StringAttribute{
							Description:         "The model of the system.",
							MarkdownDescription: "The model of the system.",
							Computed:            true,
						},
						"system_name": schema.StringAttribute{
							Description:         "Name of the system for the volume.",
							MarkdownDescription: "Name of the system for the volume.",
							Computed:            true,
						},
						"total_size": schema.Int64Attribute{
							Description:         "The total provisioned size of the volume - Unit: bytes",
							MarkdownDescription: "The total provisioned size of the volume - Unit: bytes",
							Computed:            true,
						},
						"used_size": schema.Int64Attribute{
							Description:         "The used size of the volume - Unit: bytes",
							MarkdownDescription: "The used size of the volume - Unit: bytes",
							Computed:            true,
						},
						"used_size_unique": schema.Int64Attribute{
							Description:         "The unique used size of the volume - Unit: bytes",
							MarkdownDescription: "The unique used size of the volume - Unit: bytes",
							Computed:            true,
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
	if err != nil || status.StatusCode != http.StatusOK && status.StatusCode != http.StatusPartialContent {
		errorStr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.BlockVolumesReadErrorMsg,
			errorStr,
		)
		return
	}

	// If the returned filtered values does not equal the length of the filter
	// Then one or more of the filtered values are invalid
	if filterUsed && len(volumes.Results) != len(state.Filter.IDs) {
		resp.Diagnostics.AddError(
			constants.BlockVolumesFilterErrorMsg,
			constants.FilterGeneralErrorMsg,
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
			constants.DatasourceConfigureErrorMsg,
			fmt.Sprintf(constants.GeneralConfigureErrorMsg, req.ProviderData),
		)

		return
	}

	d.client = client
}
