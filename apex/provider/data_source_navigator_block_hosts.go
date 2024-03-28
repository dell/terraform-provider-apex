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
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &hostsDataSource{}
	_ datasource.DataSourceWithConfigure = &hostsDataSource{}
)

// NewHostsDataSource returns a new hosts data source instance.
func NewHostsDataSource() datasource.DataSource {
	return &hostsDataSource{}
}

// hostsDataSource is the data source implementation.
type hostsDataSource struct {
	client *client.APIClient
}

// Metadata returns the data source type name.
func (d *hostsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_navigator_block_hosts"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (d *hostsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"hosts": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Unique Host Claim ID",
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
						"description": schema.StringAttribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"initiator_count": schema.Int64Attribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"initiator_protocol": schema.StringAttribute{
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
						"network_addresses": schema.StringAttribute{
							MarkdownDescription: " ",
							Optional:            true,
						}, "type": schema.StringAttribute{
							MarkdownDescription: " ",
							Optional:            true,
						}, "operating_system": schema.StringAttribute{
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
						"total_size": schema.Int64Attribute{
							MarkdownDescription: " ",
							Optional:            true,
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
func (d *hostsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.HostsDataSourceModel
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

	hosts, status, err := helper.GetHostCollection(d.client, filter)
	if (err != nil) || (status.StatusCode != http.StatusOK && status.StatusCode != http.StatusPartialContent) {
		resp.Diagnostics.AddError(
			"Unable to Read Apex Navigator Hosts: ",
			err.Error(),
		)
		return
	}

	// If the returned filtered values does not equal the length of the filter
	// Then one or more of the filtered values are invalid
	if filterUsed && len(hosts.Results) != len(state.Filter.IDs) {
		tflog.Debug(ctx, fmt.Sprintf("hosts: %v, filter: %v", hosts.Results, state.Filter.IDs))
		resp.Diagnostics.AddError(
			"Failed to filter hosts.",
			"one more more of the ids set in the filter is invalid.",
		)
		return
	}

	// Map response to the model
	state = helper.MapHostModel(hosts, state)
	state.ID = types.StringValue("host-ds-id")

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *hostsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
