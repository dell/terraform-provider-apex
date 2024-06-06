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
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// AwsPermissionPoliciesDataSourceSchema defines tha acceptable confirguation and state attribute names and types
var AwsPermissionPoliciesDataSourceSchema schema.Schema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Description:         "Identifier of the AWS Permission Policy",
			MarkdownDescription: "Identifier of the AWS Permission Policy",
			Optional:            true,
		},
		"version": schema.StringAttribute{
			Description:         "Version of the AWS Permission Policy",
			MarkdownDescription: "Version of the AWS Permission Policy",
			Optional:            true,
		},
		"permission_policy": schema.StringAttribute{
			MarkdownDescription: "The JSON stringified details of the Permissions Policy",
			Description:         "The JSON stringified details of the Permissions Policy",
			Computed:            true,
		},
	},
}
