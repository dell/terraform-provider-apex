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
	"github.com/hashicorp/terraform-plugin-framework/types"
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
		"permission_policy": schema.SingleNestedAttribute{
			MarkdownDescription: "Details of the Policy",
			Description:         "Details of the Policy",
			Computed:            true,
			Attributes: map[string]schema.Attribute{
				"version": schema.StringAttribute{
					Description:         "Version of the AWS Permission Policy",
					MarkdownDescription: "Version of the AWS Permission Policy",
					Optional:            true,
				},
				"statement": schema.ListNestedAttribute{
					Computed:            true,
					Description:         "The Permission Policy Statement",
					MarkdownDescription: "The Permission Policy Statement.",
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"sid": schema.StringAttribute{
								Description:         "Statement identifier of the AWS Permission Policy",
								MarkdownDescription: "Statement identifier of the AWS Permission Policy",
								Computed:            true,
							},
							"effect": schema.StringAttribute{
								Description:         "effect of the AWS Permission Policy",
								MarkdownDescription: "effect of the AWS Permission Policy",
								Computed:            true,
							},
							"action": schema.ListAttribute{
								Description:         "action of the AWS Permission Policy",
								MarkdownDescription: "action of the AWS Permission Policy",
								Computed:            true,
								ElementType:         types.StringType,
							},
							"resource": schema.StringAttribute{
								Description:         "resource of the AWS Permission Policy",
								MarkdownDescription: "resource of the AWS Permission Policy",
								Computed:            true,
							},
							"iam_aws_service_name": schema.ListAttribute{
								Description:         "IAM AWS Service Name of the AWS Permission Policy",
								MarkdownDescription: "IAM AWS Service Name of the AWS Permission Policy",
								Computed:            true,
								ElementType:         types.StringType,
							},
						},
					},
				},
			},
		},
	},
}
