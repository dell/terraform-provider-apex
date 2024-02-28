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

	"github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/dell/terraform-provider-apex/apex/models"
	client "github.com/dell/terraform-provider-apex/client/apexclient/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	resp.TypeName = req.ProviderTypeName + "_hosts"
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
	}
}

// Read method is used to refresh the Terraform state based on the schema data.
func (d *hostsDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.HostsDataSourceModel

	hosts, status, err := helper.GetHostCollection(d.client)
	if (err != nil) || (status.StatusCode != http.StatusOK && status.StatusCode != http.StatusPartialContent) {
		resp.Diagnostics.AddError(
			"Unable to Read Apex Navigator Hosts",
			err.Error(),
		)
		return
	}

	// Map response body to model
	for _, host := range hosts.Results {
		hostState := models.HostsModel{
			ID:                types.StringValue(host.Id),
			SystemID:          types.StringValue(*host.SystemId),
			SystemType:        types.StringValue(*host.SystemType),
			Description:       types.StringValue(*host.Description),
			InitiatorCount:    types.Int64Value(*host.InitiatorCount),
			InitiatorProtocol: types.StringValue(*host.InitiatorProtocol),
			IssueCount:        types.Int64Value(*host.IssueCount),
			Name:              types.StringValue(*host.Name),
			NativeID:          types.StringValue(*host.NativeId),
			NetworkAddresses:  types.StringValue(*host.NetworkAddresses),
			Type:              types.StringValue(*host.Type),
			OperatingSystem:   types.StringValue(*host.OperatingSystem),
			SystemModel:       types.StringValue(*host.SystemModel),
			SystemName:        types.StringValue(*host.SystemName),
			TotalSize:         types.Int64Value(*host.TotalSize),
		}

		state.Hosts = append(state.Hosts, hostState)
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
