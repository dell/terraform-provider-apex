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
	jmsClient "dell/apex-job-client"
	"fmt"

	"github.com/dell/terraform-provider-apex/apex/constants"
	"github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/dell/terraform-provider-apex/apex/models"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &policyGenerateResource{}
	_ resource.ResourceWithConfigure = &policyGenerateResource{}
)

// NewTrustPolicyGenerateResource returns a trust policy resource
func NewTrustPolicyGenerateResource() resource.Resource {
	return &policyGenerateResource{}
}

// policyGenerateResource is the resource implementation.
type policyGenerateResource struct {
	client     *client.APIClient
	jobsClient *jmsClient.APIClient
}

// Metadata returns the resource type name.
func (r *policyGenerateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_navigator_aws_trust_policy_generate"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (r *policyGenerateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) { // nolint:funlen, dupl

	resp.Schema = schema.Schema{
		Description:         "This Terraform resource is used to generate trust policy.",
		MarkdownDescription: "This Terraform resource is used to generate trust policy.",
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				MarkdownDescription: "The AWS account ID",
				Description:         "The AWS account ID",
				Required:            true,
			},
			"version": schema.StringAttribute{
				MarkdownDescription: "The AWS account ID",
				Description:         "The AWS account ID",
				Computed:            true,
			},
			"statement": schema.ListNestedAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Represents a list of trust policy statements",
				MarkdownDescription: "Represents a list of trust policy statements",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"effect": schema.StringAttribute{
							Description:         "Effect of the AWS Permission Policy",
							MarkdownDescription: "Effect of the AWS Permission Policy",
							Computed:            true,
							Optional:            true,
						},
						"action": schema.StringAttribute{
							Description:         "Action of the AWS Permission Policy",
							MarkdownDescription: "Action of the AWS Permission Policy",
							Computed:            true,
							Optional:            true,
						},
						"principal": schema.SingleNestedAttribute{
							Description:         "Principal of the AWS Permission Policy",
							MarkdownDescription: "Principal of the AWS Permission Policy",
							Computed:            true,
							Optional:            true,
							Attributes: map[string]schema.Attribute{
								"aws": schema.StringAttribute{
									Description:         "AWS of the AWS Permission Policy",
									MarkdownDescription: "AWS of the AWS Permission Policy",
									Computed:            true,
									Optional:            true,
								},
							},
						},
						"condition": schema.SingleNestedAttribute{
							Description:         "Condition of the AWS Permission Policy",
							MarkdownDescription: "Condition of the AWS Permission Policy",
							Computed:            true,
							Attributes: map[string]schema.Attribute{
								"string_equals": schema.SingleNestedAttribute{
									Description:         "String Equals of the AWS Permission Policy",
									MarkdownDescription: "String Equals of the AWS Permission Policy",
									Computed:            true,
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"sts_external_id": schema.StringAttribute{
											Description:         "sts_external_id of the AWS Permission Policy",
											MarkdownDescription: "sts_external_id  of the AWS Permission Policy",
											Computed:            true,
											Optional:            true,
										},
									},
								},
								"bool": schema.SingleNestedAttribute{
									Description:         "Bool of the AWS Permission Policy",
									MarkdownDescription: "Bool of the AWS Permission Policy",
									Computed:            true,
									Optional:            true,
									Attributes: map[string]schema.Attribute{
										"aws_multi_factor_auth_present": schema.StringAttribute{
											Description:         "MultiFactorAuthPresent boolean of the AWS Permission Policy",
											MarkdownDescription: "MultiFactorAuthPresent boolean of the AWS Permission Policy",
											Computed:            true,
											Optional:            true,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (r *policyGenerateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	clients, ok := req.ProviderData.(Clients)
	if !ok {
		resp.Diagnostics.AddError(
			constants.ResourceConfigureErrorMsg,
			fmt.Sprintf(constants.GeneralConfigureErrorMsg, req.ProviderData),
		)
		return
	}

	r.client = clients.APIClient
	r.jobsClient = clients.JMSClient
}

// Create creates the resource and sets the initial Terraform state.
func (r *policyGenerateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) { // nolint:funlen, gocognit

	// Get current plan
	var plan models.AwsPolicyGenerateModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// generate trust policy
	policy, status, err := helper.GenerateAwsTrustPolicy(r.client, plan.AwsAccountID.ValueString())
	if err != nil {
		newErr := helper.GetErrorString(err, status)
		resp.Diagnostics.AddError(
			constants.AwsGeneratePolicyErrorMsg,
			constants.AwsGeneratePolicyDetailMsg+newErr,
		)
		return
	}

	// Set state
	state := models.AwsPolicyGenerateModel{
		AwsAccountID: plan.AwsAccountID,
		Version:      types.StringPointerValue(policy.Version),
		Statement:    helper.MapGeneratedStatements(policy.Statement),
	}

	tflog.Info(ctx, "Created policy", map[string]interface{}{"state": state})

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read method is used to refresh the Terraform state based on the schema data.
func (r *policyGenerateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) { //nolint:dupl

	// Get current state
	var state models.AwsPolicyGenerateModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set state
	response := models.AwsPolicyGenerateModel{
		AwsAccountID: state.AwsAccountID,
		Version:      state.Version,
		Statement:    state.Statement,
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &response)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *policyGenerateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Get current state
	var state models.AwsPolicyGenerateModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current plan
	var plan models.AwsPolicyGenerateModel
	diagsPlan := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diagsPlan...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddWarning(
		constants.UpdatesNotSupportedErrorMsg,
		constants.UpdatesNotSupportedErrorMsg,
	)

	if state.AwsAccountID != plan.AwsAccountID {
		resp.Diagnostics.AddError(
			constants.AwsGeneratePolicyErrorMsg,
			constants.AwsUpdateInvalidFieldUpdateErrorMsg,
		)
	}

}

// Delete deletes the resource and removes the Terraform state on success.
func (r *policyGenerateResource) Delete(_ context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning(
		constants.DeleteIsNotSupportedErrorMsg,
		constants.AwsDeleteGeneratePolicyErrorMsg,
	)
}
