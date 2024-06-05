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

package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// AwsAccountResourceModel maps storage system schema data.
type AwsAccountResourceModel struct {
	AccountID       types.String `tfsdk:"account_id"`
	RoleArn         types.String `tfsdk:"role_arn"`
	Status          types.String `tfsdk:"status"`
	AwsAccountAlias types.String `tfsdk:"aws_account_alias"`
}

// AwsAccountDataSourceModel maps storage system schema data.
type AwsAccountDataSourceModel struct {
	Accounts []AwsAccountModel     `tfsdk:"accounts"`
	ID       types.String          `tfsdk:"id"`
	Filter   *AwsAccountFilterType `tfsdk:"filter"`
}

// AwsAccountModel maps Aws Account schema data.
type AwsAccountModel struct {
	ID     types.String `tfsdk:"id"`
	Alias  types.String `tfsdk:"account_alias"`
	Status types.String `tfsdk:"status"`
}

// AwsAccountFilterType describes the filter data model.
type AwsAccountFilterType struct {
	IDs []types.String `tfsdk:"ids"`
}

// AwsPermissionsDataSourceModel maps the data source schema data.
type AwsPermissionsDataSourceModel struct {
	Permissions []AwsPermissionsModel     `tfsdk:"permissions"`
	ID          types.String              `tfsdk:"id"`
	Filter      *AwsPermissionsFilterType `tfsdk:"filter"`
}

// AwsPermissionsFilterType describes the filter data model.
type AwsPermissionsFilterType struct {
	IDs []types.String `tfsdk:"ids"`
}

// AwsPermissionsModel maps the Aws Permissions schema data.
type AwsPermissionsModel struct {
	ID      types.String `tfsdk:"id"`
	Version types.String `tfsdk:"version"`
	// This is the JSON Stringified version of the permissions object
	PermissionsPolicy types.String `tfsdk:"permission_policy"`
}

// Statement describes the statement data model.
type Statement struct {
	Sid               types.String   `tfsdk:"sid"`
	Effect            types.String   `tfsdk:"effect"`
	Action            []types.String `tfsdk:"action"`
	Resource          types.String   `tfsdk:"resource"`
	IamAwsServiceName []types.String `tfsdk:"iam_aws_service_name"`
}

// AwsPolicyGenerateModel maps the Aws Permissions schema data.
type AwsPolicyGenerateModel struct {
	AwsAccountID types.String        `tfsdk:"account_id"`
	Version      types.String        `tfsdk:"version"`
	Statement    basetypes.ListValue `tfsdk:"statement"`
}

// PolicyGenerateStatement describes the statement data model.
type PolicyGenerateStatement struct {
	Effect    types.String    `tfsdk:"effect"`
	Action    types.String    `tfsdk:"action"`
	Principal PolicyPrincipal `tfsdk:"principal"`
	Condition Condition       `tfsdk:"condition"`
}

// PolicyPrincipal describes the principal data model.
type PolicyPrincipal struct {
	AWS types.String `tfsdk:"aws"`
}

// Condition describes the condition data model.
type Condition struct {
	Bool         BoolStatment `tfsdk:"bool"`
	StringEquals StringEquals `tfsdk:"string_equals"`
}

// BoolStatment describes the bool data model.
type BoolStatment struct {
	AwsMultiFactorAuthPresent types.String `tfsdk:"aws_multi_factor_auth_present"`
}

// StringEquals describes the string_equals data model.
type StringEquals struct {
	StsExternalID types.String `tfsdk:"sts_external_id"`
}
