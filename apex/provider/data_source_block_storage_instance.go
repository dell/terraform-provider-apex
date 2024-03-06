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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &blockStorageInstanceDataSource{}
	_ datasource.DataSourceWithConfigure = &blockStorageInstanceDataSource{}
)

// NewBlockStorageInstanceDataSource is a block storage data source object
func NewBlockStorageInstanceDataSource() datasource.DataSource {
	return &blockStorageInstanceDataSource{}
}

// storageProductsDataSource is the data source implementation.
type blockStorageInstanceDataSource struct {
	client *client.APIClient
}

func (d *blockStorageInstanceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_block_storage"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (d *blockStorageInstanceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) { // nolint:funlen
	resp.Schema = BlockStorageDataSourceSchema
}

// Read method is used to refresh the Terraform state based on the schema data.
func (d *blockStorageInstanceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) { // nolint:gocognit, funlen
	var state models.BlockStorageModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	storageSystem, status, err := helper.GetBlockStorageInstance(d.client, state.ID.ValueString())
	if (err != nil) || (status.StatusCode != http.StatusOK) {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Unable to Read Apex Navigator Block Storage"+newErr,
			err.Error(),
		)
		return
	}

	// needs nested for loop for deployment details
	// Map response body to model
	storageSystemsState := helper.GetBlockStorageSystem(*storageSystem)
	if storageSystem.DeploymentDetails != nil {
		if storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails != nil {
			storageSystemsState.DeploymentDetails.SystemPublicCloud.VirtualPrivateCloud = types.StringPointerValue(storageSystem.DeploymentDetails.SystemPublicCloudDeploymentDetails.VirtualPrivateCloud)
		}
	} else {
		resp.Diagnostics.AddError(
			"Unexpected system type",
			"Unexpected system type",
		)
	}

	// Set state
	diags := resp.State.Set(ctx, &storageSystemsState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *blockStorageInstanceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
