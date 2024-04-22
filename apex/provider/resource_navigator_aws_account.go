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
	jmsClient "dell/apex-job-client"
	"fmt"
	"net/http"

	client "dell/apex-client"

	"github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/dell/terraform-provider-apex/apex/models"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &clonesResource{}
	_ resource.ResourceWithConfigure   = &clonesResource{}
	_ resource.ResourceWithImportState = &clonesResource{}
)

// NewAwsAccountResource is a helper function to simplify the provider implementation.
func NewAwsAccountResource() resource.Resource {
	return &awsAccountResource{}
}

// awsAccountResource is the resource implementation.
type awsAccountResource struct {
	client     *client.APIClient
	jobsClient *jmsClient.APIClient
}

// Metadata returns the resource type name.
func (r *awsAccountResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_navigator_aws_account"
}

// Schema defines the schema for the resource.
func (r *awsAccountResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This Terraform resource is used to manage aws accounts on Apex Navigator. We can connect, read, update aws role and disconnect the aws accounts using this resource. We can also import an existing aws account from Apex Navigator.",
		MarkdownDescription: "This Terraform resource is used to manage aws accounts on Apex Navigator. We can connect, read, update aws role and disconnect the aws accounts using this resource. We can also import an existing aws account from Apex Navigator.",
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the AWS account, this value can not be updated after connection.",
				Description:         "The ID of the AWS account, this value can not be updated after connection.",
				Required:            true,
			},
			"role_arn": schema.StringAttribute{
				MarkdownDescription: "The role ARN of the AWS account",
				Description:         "The role ARN of the AWS account",
				Required:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "The status of the AWS account",
				Description:         "The status of the AWS account",
				Computed:            true,
				Optional:            true,
			},
			"aws_account_alias": schema.StringAttribute{
				MarkdownDescription: "The alias of the AWS account",
				Description:         "The alias of the AWS account",
				Computed:            true,
				Optional:            true,
			},
		},
	}
}

// Configure adds the provider configured client to the resource.
func (r *awsAccountResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	clients, ok := req.ProviderData.(Clients)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *provider.CLients, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = clients.APIClient
	r.jobsClient = clients.JMSClient
}

// Create creates the resource and sets the initial Terraform state.
func (r *awsAccountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan.
	var plan models.AwsAccountResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the resource
	accountPost := *client.NewAwsAccountsCreateInput(plan.AccountID.ValueString())

	account, status, err := helper.ConnectAccount(r.client, accountPost)
	if err != nil || status == nil || status.StatusCode != http.StatusAccepted {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error connecting AWS Account",
			"Could not connect AWS Account, unexpected error: "+newErr,
		)
		return
	}

	plan.Status = types.StringPointerValue(account.Status)
	plan.AwsAccountAlias = types.StringPointerValue(account.AwsAccountAlias)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *awsAccountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Retrieve values from state.
	var state models.AwsAccountResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed Terraform state
	account, status, err := helper.GetAccountInstance(r.client, state.AccountID.ValueString())
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error Reading Apex Navigator Aws Accounts",
			"Could not read AWS Accounts, unexpected error: "+newErr,
		)
		return
	}
	state.Status = types.StringPointerValue(account.Status)
	state.AwsAccountAlias = types.StringPointerValue(account.AwsAccountAlias)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *awsAccountResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan models.AwsAccountResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieve values from state
	var state models.AwsAccountResourceModel
	diagsState := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diagsState...)
	if resp.Diagnostics.HasError() {
		return
	}

	if state.AccountID != plan.AccountID {
		resp.Diagnostics.AddError(
			"Error Updating Apex Navigator Aws Accounts",
			"Account ID cannot be updated",
		)
		return
	}

	// Update the account
	updateJob, status, err := helper.UpdateAccount(r.client, plan.AccountID.ValueString(), plan.RoleArn.ValueString())
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error Updating Apex Navigator Aws Accounts",
			"Could not update AWS Accounts, unexpected error: "+newErr,
		)
		return
	}

	// Fetching Job Status
	_, jobErr := helper.WaitForJobToComplete(ctx, r.jobsClient, updateJob.Id)
	if jobErr != nil {
		resp.Diagnostics.AddError(
			"Error occurred during update job: ",
			helper.ResourceRetrieveError+jobErr.Error(),
		)
		return
	}

	// Get refreshed Terraform state
	account, status, err := helper.GetAccountInstance(r.client, plan.AccountID.ValueString())
	if err != nil || status == nil || status.StatusCode != http.StatusOK {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error Reading Apex Navigator AWS Accounts",
			"Could not read AWS Accounts, unexpected error: "+newErr,
		)
		return
	}
	plan.Status = types.StringPointerValue(account.Status)
	plan.AwsAccountAlias = types.StringPointerValue(account.AwsAccountAlias)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *awsAccountResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state models.AwsAccountResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// disconnect the account
	_, status, err := helper.DisconnectAccount(r.client, state.AccountID.ValueString())
	if err != nil || status == nil {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			"Error Deleting Apex Navigator Aws Accounts",
			"Could not delete AWS Accounts, unexpected error: "+newErr,
		)
		return
	}

}

// ImportsState imports the Terraform resource
func (r *awsAccountResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("account_id"), req, resp)
}
