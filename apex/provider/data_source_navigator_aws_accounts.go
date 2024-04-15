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
	client "dell/apex-client"
	"fmt"
	"net/http"

	"github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/dell/terraform-provider-apex/apex/models"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &awsAccountsDataSource{}
	_ datasource.DataSourceWithConfigure = &awsAccountsDataSource{}
)

// NewAwsAccountsDataSource returns a new aws accounts datasource instance.
func NewAwsAccountsDataSource() datasource.DataSource {
	return &awsAccountsDataSource{}
}

// awsAccountsDataSource is the data source implementation.
type awsAccountsDataSource struct {
	client *client.APIClient
}

// Metadata returns the data source type name.
func (d *awsAccountsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_navigator_aws_accounts"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (d *awsAccountsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "This Terraform Datasource is used to query existing aws accounts on Apex Navigator." +
			" The information fetched from this block can be further used for resource block.",
		MarkdownDescription: "This Terraform Datasource is used to query existing aws accounts on Apex Navigator." +
			" The information fetched from this block can be further used for resource block.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "ID of the datasource",
				MarkdownDescription: "ID of the datasource",
				Computed:            true,
			},
			"accounts": schema.ListNestedAttribute{
				Description:         "List of aws accounts",
				MarkdownDescription: "List of aws accounts",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: AwsAccountsDataSourceSchema.Attributes,
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
func (d *awsAccountsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.AwsAccountDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if (state.Filter == nil) || (state.Filter.IDs == nil) {
		state.Filter = &models.AwsAccountFilterType{
			IDs: []types.String{},
		}
	}

	accounts, status, err := helper.GetAwsAccountCollection(d.client, state.Filter.IDs)
	if (err != nil) || (status.StatusCode != http.StatusOK && status.StatusCode != http.StatusPartialContent) {
		resp.Diagnostics.AddError(
			"Unable to Read Apex Navigator Aws Accounts",
			err.Error(),
		)
		return
	}

	// If the returned filtered values does not equal the length of the filter
	// Then one or more of the filtered values are invalid
	if len(state.Filter.IDs) > 0 && len(accounts.Results) != len(state.Filter.IDs) {
		resp.Diagnostics.AddError(
			"Failed to filter aws accounts.",
			"one more more of the ids set in the filter is invalid.",
		)
		return
	}

	// Map response body to model
	for _, account := range accounts.Results {
		accountState := helper.MapAwsAccount(account)
		state.Accounts = append(state.Accounts, accountState)
	}
	state.ID = types.StringValue("aws-account-ds-id")

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

// Configure adds the provider configured client to the data source.
func (d *awsAccountsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
