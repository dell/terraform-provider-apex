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

// ClonesModel defines the attribute names and types for a Clones TF model
type ClonesModel struct {
	ID                types.String        `tfsdk:"id"`
	Name              types.String        `tfsdk:"name"`
	Description       types.String        `tfsdk:"description"`
	MobilityTargetID  types.String        `tfsdk:"mobility_target_id"`
	CreationTimestamp types.String        `tfsdk:"creation_timestamp"`
	RefreshTimestamp  types.String        `tfsdk:"refresh_timestamp"`
	ImageTimestamp    types.String        `tfsdk:"image_timestamp"`
	CloneVolumes      basetypes.ListValue `tfsdk:"clone_volumes"`
	HostMappings      basetypes.ListValue `tfsdk:"host_mappings"`
}

// ClonesDataSourceModel maps the data source schema data.
type ClonesDataSourceModel struct {
	Clones []ClonesModel    `tfsdk:"clones"`
	ID     types.String     `tfsdk:"id"`
	Filter *CloneFilterType `tfsdk:"filter"`
}

// CloneFilterType describes the filter data model.
type CloneFilterType struct {
	IDs []types.String `tfsdk:"ids"`
}

// ClonesRefreshModel defines the attribute names and types for a Clones Refresh TF model
type ClonesRefreshModel struct {
	ID      types.String `tfsdk:"id"`
	CloneID types.String `tfsdk:"clone_id"`
	Status  types.String `tfsdk:"status"`
}
