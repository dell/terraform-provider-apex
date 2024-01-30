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

package navigator

import (
	"context"
	"fmt"
	"net/http"

	client "github.com/dell/terraform-provider-apex/client/apexclient/client"
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

// volumesDataSourceModel maps the data source schema data.
type volumesDataSourceModel struct {
	Volumes []volumesModel `tfsdk:"volumes"`
	ID      types.String   `tfsdk:"id"`
}

// volumesModel maps coffees schema data.
type volumesModel struct {
	ID                      types.String  `tfsdk:"id"`
	SystemID                types.String  `tfsdk:"system_id"`
	SystemType              types.String  `tfsdk:"system_type"`
	AllocatedSize           types.Int64   `tfsdk:"allocated_size"`
	Bandwidth               types.Int64   `tfsdk:"bandwidth"`
	ConsistencyGroupName    types.String  `tfsdk:"consistency_group_name"`
	DataReductionPercent    types.Float64 `tfsdk:"data_reduction_percent"`
	DataReductionRatio      types.Float64 `tfsdk:"data_reduction_ratio"`
	DataReductionSavedSize  types.Int64   `tfsdk:"data_reduction_saved_size"`
	IoLimitPolicyName       types.String  `tfsdk:"io_limit_policy_name"`
	Iops                    types.Int64   `tfsdk:"iops"`
	IsCompressedOrDeduped   types.String  `tfsdk:"is_compressed_or_deduped"`
	IsThinEnabled           types.Bool    `tfsdk:"is_thin_enabled"`
	IssueCount              types.Int64   `tfsdk:"issue_count"`
	Latency                 types.Int64   `tfsdk:"latency"`
	LogicalSize             types.Int64   `tfsdk:"logical_size"`
	Name                    types.String  `tfsdk:"name"`
	NativeID                types.String  `tfsdk:"native_id"`
	Type                    types.String  `tfsdk:"type"`
	PoolID                  types.String  `tfsdk:"pool_id"`
	PoolName                types.String  `tfsdk:"pool_name"`
	PoolType                types.String  `tfsdk:"pool_type"`
	SnapshotCount           types.Int64   `tfsdk:"snap_shot_count"`
	SnapshotPolicy          types.String  `tfsdk:"snap_shot_policy"`
	SnapshotSize            types.Int64   `tfsdk:"snap_shot_size"`
	StorageResourceID       types.String  `tfsdk:"storage_resource_id"`
	StorageResourceNativeID types.String  `tfsdk:"storage_resource_native_id"`
	SystemModel             types.String  `tfsdk:"system_model"`
	SystemName              types.String  `tfsdk:"system_name"`
	TotalSize               types.Int64   `tfsdk:"total_size"`
	UsedSize                types.Int64   `tfsdk:"used_size"`
	UsedSizeUnique          types.Int64   `tfsdk:"used_size_unique"`
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
	}
}

// Read method is used to refresh the Terraform state based on the schema data.
func (d *volumesDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state volumesDataSourceModel

	volumes, status, err := d.client.VolumesAPI.VolumesCollection(context.Background()).Limit(500).Execute()
	if (err != nil) || (status.StatusCode != http.StatusOK && status.StatusCode != http.StatusPartialContent) {
		resp.Diagnostics.AddError(
			"Unable to Read Apex Navigator Volumes",
			err.Error(),
		)
		return
	}
	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}

	// Map response body to model
	for _, volume := range volumes.GetResults() {
		volumeState := volumesModel{
			ID:                      types.StringValue(volume.Id),
			SystemID:                types.StringValue(*volume.SystemId),
			SystemType:              types.StringValue(*volume.SystemType),
			AllocatedSize:           types.Int64Value(*volume.AllocatedSize),
			Bandwidth:               types.Int64Value(*volume.Bandwidth),
			ConsistencyGroupName:    types.StringValue(*volume.ConsistencyGroupName),
			DataReductionPercent:    types.Float64Value(*volume.DataReductionPercent),
			DataReductionRatio:      types.Float64Value(*volume.DataReductionRatio),
			DataReductionSavedSize:  types.Int64Value(*volume.DataReductionSavedSize),
			IoLimitPolicyName:       types.StringValue(*volume.IoLimitPolicyName),
			Iops:                    types.Int64Value(*volume.Iops),
			IsCompressedOrDeduped:   types.StringValue(*volume.IsCompressedOrDeduped),
			IsThinEnabled:           types.BoolValue(*volume.IsThinEnabled),
			IssueCount:              types.Int64Value(*volume.IssueCount),
			Latency:                 types.Int64Value(*volume.Latency),
			LogicalSize:             types.Int64Value(*volume.LogicalSize),
			Name:                    types.StringValue(*volume.Name),
			NativeID:                types.StringValue(*volume.NativeId),
			Type:                    types.StringValue(*volume.Type),
			PoolID:                  types.StringValue(*volume.PoolId),
			PoolName:                types.StringValue(*volume.PoolName),
			PoolType:                types.StringValue(*volume.PoolType),
			SnapshotCount:           types.Int64Value(*volume.SnapshotCount),
			SnapshotPolicy:          types.StringValue(*volume.SnapshotPolicy),
			SnapshotSize:            types.Int64Value(*volume.SnapshotSize),
			StorageResourceID:       types.StringValue(*volume.StorageResourceId),
			StorageResourceNativeID: types.StringValue(*volume.StorageResourceNativeId),
			SystemModel:             types.StringValue(*volume.SystemModel),
			SystemName:              types.StringValue(*volume.SystemName),
			TotalSize:               types.Int64Value(*volume.TotalSize),
			UsedSize:                types.Int64Value(*volume.UsedSize),
			UsedSizeUnique:          types.Int64Value(*volume.UsedSizeUnique),
		}
		state.Volumes = append(state.Volumes, volumeState)
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
