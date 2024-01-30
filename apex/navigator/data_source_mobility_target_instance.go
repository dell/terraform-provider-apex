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

package navigator //nolint:dupl

import (
	"context"
	"fmt"
	"net/http"

	client "github.com/dell/terraform-provider-apex/client/apex"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &mobilityTargetInstanceDataSource{}
	_ datasource.DataSourceWithConfigure = &mobilityTargetInstanceDataSource{}
)

// NewMobilityTargetInstanceDataSource is a storage system data source object
func NewMobilityTargetInstanceDataSource() datasource.DataSource {
	return &mobilityTargetInstanceDataSource{}
}

// mobilityTargetInstanceDataSource is the data source implementation.
type mobilityTargetInstanceDataSource struct {
	client *client.APIClient
}

func (d *mobilityTargetInstanceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mobility_target"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (d *mobilityTargetInstanceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) { // nolint:funlen, dupl
	resp.Schema = MobilityTargetsDataSourceSchema
}

// Read method is used to refresh the Terraform state based on the schema data.
func (d *mobilityTargetInstanceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) { // nolint:gocognit, funlen
	var state MobilityTargetModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	mobilityTarget, status, err := d.client.MobilityTargetsAPI.MobilityTargetsInstance(context.Background(), state.ID.ValueString()).Execute()
	if (err != nil) || (status.StatusCode != http.StatusOK) {
		resp.Diagnostics.AddError(
			"Unable to Read Apex Navigator Mobility Target",
			err.Error(),
		)
		return
	}

	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}

	// Map response body to model
	mobilityTargetsState := GetMobilityTargetModel(*mobilityTarget)

	state.ID = types.StringValue("placeholder")

	// Set state
	diags := resp.State.Set(ctx, &mobilityTargetsState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *mobilityTargetInstanceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
