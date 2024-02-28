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

// HostsDataSourceModel maps the data source schema data.
type HostsDataSourceModel struct {
	Hosts []HostsModel `tfsdk:"hosts"`
	ID    types.String `tfsdk:"id"`
}

// HostsModel maps hosts schema data.
type HostsModel struct {
	ID                types.String `tfsdk:"id"`
	SystemID          types.String `tfsdk:"system_id"`
	SystemType        types.String `tfsdk:"system_type"`
	Description       types.String `tfsdk:"description"`
	InitiatorCount    types.Int64  `tfsdk:"initiator_count"`
	InitiatorProtocol types.String `tfsdk:"initiator_protocol"`
	IssueCount        types.Int64  `tfsdk:"issue_count"`
	Name              types.String `tfsdk:"name"`
	NativeID          types.String `tfsdk:"native_id"`
	NetworkAddresses  types.String `tfsdk:"network_addresses"`
	Type              types.String `tfsdk:"type"`
	OperatingSystem   types.String `tfsdk:"operating_system"`
	SystemModel       types.String `tfsdk:"system_model"`
	SystemName        types.String `tfsdk:"system_name"`
	TotalSize         types.Int64  `tfsdk:"total_size"`
}
