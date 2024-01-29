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

package provider //nolint:dupl

import (
	"context"
	"fmt"
	"net/http"

	"eos2git.cec.lab.emc.com/CIRRUS/cirrus-terraform-providers/pkg/gen/apex/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &clonesInstanceDataSource{}
	_ datasource.DataSourceWithConfigure = &clonesInstanceDataSource{}
)

// NewCloneInstancesDataSource returns a new clone data source instance.
func NewCloneInstancesDataSource() datasource.DataSource {
	return &clonesInstanceDataSource{}
}

// clonesInstanceDataSource is the data source implementation.
type clonesInstanceDataSource struct {
	client *client.APIClient
}

func (d *clonesInstanceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clone"
}

func (d *clonesInstanceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ClonesDataSourceSchema
}
func (d *clonesInstanceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ClonesModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	clones, status, err := d.client.ClonesAPI.ClonesInstance(context.Background(), state.ID.ValueString()).Execute()
	if (err != nil) || (status.StatusCode != http.StatusOK) {
		resp.Diagnostics.AddError(
			"Unable to Read Apex Navigator Clones",
			err.Error(),
		)
		return
	}

	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}

	// Map response body to model
	clonesState := GetClonesModel(*clones)

	state.ID = types.StringValue("placeholder")

	// Set state
	diags := resp.State.Set(ctx, &clonesState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *clonesInstanceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
