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
)

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
	ID                types.String      `tfsdk:"id"`
	Version           types.String      `tfsdk:"version"`
	PermissionsPolicy PermissionsPolicy `tfsdk:"permission_policy"`
}

// PermissionsPolicy describes the permissions policy data model.
type PermissionsPolicy struct {
	Version   types.String `tfsdk:"version"`
	Statement []Statement  `tfsdk:"statement"`
}

// Statement describes the statement data model.
type Statement struct {
	Sid               types.String   `tfsdk:"sid"`
	Effect            types.String   `tfsdk:"effect"`
	Action            []types.String `tfsdk:"action"`
	Resource          types.String   `tfsdk:"resource"`
	IamAwsServiceName []types.String `tfsdk:"iam_aws_service_name"`
}
