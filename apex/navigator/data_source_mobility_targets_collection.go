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
	_ datasource.DataSource              = &mobilityTargetsDataSource{}
	_ datasource.DataSourceWithConfigure = &mobilityTargetsDataSource{}
)

// NewMobilityTargetsDataSource is a storage system data source object
func NewMobilityTargetsDataSource() datasource.DataSource {
	return &mobilityTargetsDataSource{}
}

// mobilityTargetsDataSource is the data source implementation.
type mobilityTargetsDataSource struct {
	client *client.APIClient
}

type mobilityTargetsDataSourceModel struct {
	MobilityTargets []MobilityTargetModel `tfsdk:"mobility_targets"`
	ID              types.String          `tfsdk:"id"`
}

func (d *mobilityTargetsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mobility_targets"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (d *mobilityTargetsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) { // nolint:funlen, dupl
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"mobility_targets": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: MobilityTargetsDataSourceSchema.Attributes,
				},
			},
		},
	}
}

// Read method is used to refresh the Terraform state based on the schema data.
func (d *mobilityTargetsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) { // nolint:gocognit, funlen
	var state mobilityTargetsDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	mobilityTargets, status, err := d.client.MobilityTargetsAPI.MobilityTargetsCollection(context.Background()).Limit(500).Execute()
	if (err != nil) || (status.StatusCode != http.StatusOK && status.StatusCode != http.StatusPartialContent) {
		resp.Diagnostics.AddError(
			"Unable to Read Apex Navigator Mobility Targets",
			err.Error(),
		)
		return
	}

	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}

	// Map response body to model
	for _, mobilityTarget := range mobilityTargets.Results {
		mobilityTargetsState := GetMobilityTargetModel(mobilityTarget)

		state.MobilityTargets = append(state.MobilityTargets, mobilityTargetsState)
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
func (d *mobilityTargetsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
