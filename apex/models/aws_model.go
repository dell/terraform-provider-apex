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

// AwsAccountModel maps storage system schema data.
type AwsAccountModel struct {
	ID     types.String `tfsdk:"id"`
	Alias  types.String `tfsdk:"account_alias"`
	Status types.String `tfsdk:"status"`
}

// AwsAccountFilterType describes the filter data model.
type AwsAccountFilterType struct {
	IDs []types.String `tfsdk:"ids"`
}
