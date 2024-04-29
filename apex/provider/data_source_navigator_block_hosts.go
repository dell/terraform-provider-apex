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
		Description: "This Terraform Datasource is used to query existing hosts on Apex Navigator." +
			" The information fetched from this block can be further used for resource block.",
		MarkdownDescription: "This Terraform Datasource is used to query existing hosts on Apex Navigator." +
			" The information fetched from this block can be further used for resource block.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "ID of the host datasource",
				MarkdownDescription: "ID of the host datasource",
				Computed:            true,
			},
			"hosts": schema.ListNestedAttribute{
				Computed:            true,
				MarkdownDescription: "List of hosts",
				Description:         "List of hosts",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Host identifier",
							Optional:            true,
						},
						"system_id": schema.StringAttribute{
							MarkdownDescription: "Unique identifier for the system that the host is connected to",
							Description:         "Unique identifier for the system that the host is connected to",
							Optional:            true,
						},
						"system_type": schema.StringAttribute{
							Description:         "Product type of the system",
							MarkdownDescription: "Product type of the system",
							Optional:            true,
						},
						"description": schema.StringAttribute{
							Description:         "Description of the host",
							MarkdownDescription: "Description of the host",
							Optional:            true,
						},
						"initiator_count": schema.Int64Attribute{
							Description:         "Number of initiators that are connected between the host or server and the monitored system",
							MarkdownDescription: "Number of initiators that are connected between the host or server and the monitored system",
							Optional:            true,
						},
						"initiator_protocol": schema.StringAttribute{
							Description:         "Type of initiator (FC or iSCSI) that the host or server uses to connect to a monitored system",
							MarkdownDescription: "Type of initiator (FC or iSCSI) that the host or server uses to connect to a monitored system",
							Optional:            true,
						},
						"issue_count": schema.Int64Attribute{
							Description:         "Number of health issues that are present on the host or server",
							MarkdownDescription: "Number of health issues that are present on the host or server",
							Optional:            true,
						},
						"name": schema.StringAttribute{
							Description:         "Name of the host or server",
							MarkdownDescription: "Name of the host or server",
							Optional:            true,
						},
						"native_id": schema.StringAttribute{
							Description:         "Identifier of the host, defined by the system",
							MarkdownDescription: "Identifier of the host, defined by the system",
							Optional:            true,
						},
						"network_addresses": schema.StringAttribute{
							Description:         "IPv4 or IPv6 IP addresses, domain names, or netgroup name associated with the host or server",
							MarkdownDescription: " ",
							Optional:            true,
						},
						"type": schema.StringAttribute{
							Description:         "Type of the host",
							MarkdownDescription: "Type of the host",
							Optional:            true,
						},
						"operating_system": schema.StringAttribute{
							Description:         "Operating system of the host or server",
							MarkdownDescription: "Operating system of the host or server",
							Optional:            true,
						},
						"system_model": schema.StringAttribute{
							Description:         "Model of the system",
							MarkdownDescription: "Model of the system",
							Optional:            true,
						},
						"system_name": schema.StringAttribute{
							Description:         "Name of the system",
							MarkdownDescription: "Name of the system",
							Optional:            true,
						},
						"total_size": schema.Int64Attribute{
							Description:         "Total size of all LUNs or Volumes that are provisioned to the host or server from the system - Unit: bytes",
							MarkdownDescription: "Total size of all LUNs or Volumes that are provisioned to the host or server from the system - Unit: bytes",
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
			constants.BlockHostsReadErrorMsg,
			err.Error(),
		)
		return
	}

	// If the returned filtered values does not equal the length of the filter
	// Then one or more of the filtered values are invalid
	if filterUsed && len(hosts.Results) != len(state.Filter.IDs) {
		tflog.Debug(ctx, fmt.Sprintf("hosts: %v, filter: %v", hosts.Results, state.Filter.IDs))
		resp.Diagnostics.AddError(
			constants.FilterGeneralErrorMsg,
			constants.FilterGeneralErrorMsg,
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
			constants.DatasourceConfigureErrorMsg,
			fmt.Sprintf(constants.GeneralConfigureErrorMsg, req.ProviderData),
		)

		return
	}

	d.client = client
}
