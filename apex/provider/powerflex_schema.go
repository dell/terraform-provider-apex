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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
)

// PowerflexInfo Schema defines tha acceptable confirguation and state attribute names and types
var PowerflexInfo schema.Schema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"username": schema.StringAttribute{
			MarkdownDescription: "Username of the powerflex",
			Description:         "Username of the powerflex",
			Required:            true,
		},
		"password": schema.StringAttribute{
			MarkdownDescription: "Password of the powerflex",
			Description:         "Password of the powerflex",
			Sensitive:           true,
			Required:            true,
		},
		"insecure": schema.BoolAttribute{
			MarkdownDescription: "Validated the certificate when connecting to the powerflex, defaults if unset to true",
			Description:         "Validated the certificate when connecting to the powerflex, defaults if unset to true",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(true),
		},
		"scheme": schema.StringAttribute{
			MarkdownDescription: "Scheme of the powerflex, defaults if unset to https",
			Description:         "Scheme of the powerflex, defaults if unset to https",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("https"),
		},
		"host": schema.StringAttribute{
			MarkdownDescription: "Host, ip or hostname of the powerflex. If left empty we will attempt to get the ip through Apex from the ID",
			Description:         "Host, ip or hostname of the powerflex. If left empty we will attempt to get the ip through Apex from the ID",
			Optional:            true,
			Computed:            true,
		},
	},
}

// PowerflexAsyncInfo Schema defines tha acceptable confirguation and state attribute names and types
var PowerflexAsyncInfo schema.Schema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"username": schema.StringAttribute{
			MarkdownDescription: "Username of the powerflex",
			Description:         "Username of the powerflex",
			Required:            true,
		},
		"password": schema.StringAttribute{
			MarkdownDescription: "Password of the powerflex",
			Description:         "Password of the powerflex",
			Sensitive:           true,
			Required:            true,
		},
		"insecure": schema.BoolAttribute{
			MarkdownDescription: "Validated the certificate when connecting to the powerflex, defaults if unset to true",
			Description:         "Validated the certificate when connecting to the powerflex, defaults if unset to true",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(true),
		},
		"scheme": schema.StringAttribute{
			MarkdownDescription: "Scheme of the powerflex, defaults if unset to https",
			Description:         "Scheme of the powerflex, defaults if unset to https",
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("https"),
		},
		"host": schema.StringAttribute{
			MarkdownDescription: "Host, ip or hostname of the powerflex. If left empty we will attempt to get the ip through Apex from the ID",
			Description:         "Host, ip or hostname of the powerflex. If left empty we will attempt to get the ip through Apex from the ID",
			Optional:            true,
			Computed:            true,
		},
		"poll_interval": schema.Int64Attribute{
			MarkdownDescription: "Poll interval for long-running operations in seconds. This will check if the powerflex is activated every interval. Defaults if unset to 30",
			Description:         "Poll interval for long-running operations in seconds. This will check if the powerflex is activated every interval. Defaults if unset to 30",
			Optional:            true,
			Computed:            true,
			Default:             int64default.StaticInt64(30),
		},
	},
}
