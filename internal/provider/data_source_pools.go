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

	"github.com/dell/terraform-provider-aex/pkg/gen/apex/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &poolsDataSource{}
	_ datasource.DataSourceWithConfigure = &poolsDataSource{}
)

// NewPoolsDataSource returns a new pools data source instance.
func NewPoolsDataSource() datasource.DataSource {
	return &poolsDataSource{}
}

// poolsDataSource is the data source implementation.
type poolsDataSource struct {
	client *client.APIClient
}

// poolsDataSourceModel maps the data source schema data.
type poolsDataSourceModel struct {
	Pools []poolsModel `tfsdk:"pools"`
	ID    types.String `tfsdk:"id"`
}

// poolsModel maps pools schema data.
type poolsModel struct {
	ID                   types.String  `tfsdk:"id"`
	SystemID             types.String  `tfsdk:"system_id"`
	SystemType           types.String  `tfsdk:"system_type"`
	FreeSize             types.Int64   `tfsdk:"free_size"`
	IssueCount           types.Int64   `tfsdk:"issue_count"`
	Name                 types.String  `tfsdk:"name"`
	NativeID             types.String  `tfsdk:"native_id"`
	SubscribedPercent    types.Float64 `tfsdk:"subscribed_percent"`
	SubscribedSize       types.Int64   `tfsdk:"subscribed_size"`
	SystemModel          types.String  `tfsdk:"system_model"`
	SystemName           types.String  `tfsdk:"system_name"`
	TimeToFullPrediction types.String  `tfsdk:"time_to_full_prediction"`
	TotalSize            types.Int64   `tfsdk:"total_size"`
	Type                 types.String  `tfsdk:"type"`
	UsedPercent          types.Float64 `tfsdk:"used_percent"`
	UsedSize             types.Int64   `tfsdk:"used_size"`
}

// Metadata returns the data source type name.
func (d *poolsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pools"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (d *poolsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"pools": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Unique Pool Claim ID",
							Optional:            true,
						},
						"system_id": schema.StringAttribute{
							MarkdownDescription: "",
							Optional:            true,
						},
						"system_type": schema.StringAttribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"free_size": schema.Int64Attribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"issue_count": schema.Int64Attribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"native_id": schema.StringAttribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"subscribed_percent": schema.Float64Attribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"subscribed_size": schema.Int64Attribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"system_model": schema.StringAttribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"system_name": schema.StringAttribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"time_to_full_prediction": schema.StringAttribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"total_size": schema.Int64Attribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"used_percent": schema.Float64Attribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"used_size": schema.Int64Attribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

// Read method is used to refresh the Terraform state based on the schema data.
func (d *poolsDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state poolsDataSourceModel

	pools, status, err := d.client.PoolsAPI.PoolsCollection(context.Background()).Limit(500).Execute()
	if (err != nil) || (status.StatusCode != http.StatusOK) {
		resp.Diagnostics.AddError(
			"Unable to Read Apex Navigator Pools",
			err.Error(),
		)
		return
	}

	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}

	// Map response body to model
	for _, pool := range pools.Results {
		poolState := poolsModel{
			ID:                   types.StringValue(pool.Id),
			SystemID:             types.StringValue(*pool.SystemId),
			SystemType:           types.StringValue(*pool.SystemType),
			FreeSize:             types.Int64Value(*pool.FreeSize),
			IssueCount:           types.Int64Value(*pool.IssueCount),
			Name:                 types.StringValue(*pool.Name),
			NativeID:             types.StringValue(*pool.NativeId),
			SubscribedPercent:    types.Float64Value(*pool.SubscribedPercent),
			SubscribedSize:       types.Int64Value(*pool.SubscribedSize),
			SystemModel:          types.StringValue(*pool.SystemModel),
			SystemName:           types.StringValue(*pool.SystemName),
			TimeToFullPrediction: types.StringValue(*pool.TimeToFullPrediction),
			TotalSize:            types.Int64Value(*pool.TotalSize),
			Type:                 types.StringValue(*pool.Type),
			UsedPercent:          types.Float64Value(*pool.UsedPercent),
			UsedSize:             types.Int64Value(*pool.UsedSize),
		}

		state.Pools = append(state.Pools, poolState)
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
func (d *poolsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
