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

import "github.com/hashicorp/terraform-plugin-framework/types"

// ActivationClientModel defines the attribute names and types for a PowerFlex Client TF model
type ActivationClientModel struct {
	Host     types.String `tfsdk:"host"`
	Insecure types.Bool   `tfsdk:"insecure"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
	Scheme   types.String `tfsdk:"scheme"`
}

// ActivationAsyncClientModel defines the attribute names and types for a PowerFlex Client TF model
type ActivationAsyncClientModel struct {
	Host         types.String `tfsdk:"host"`
	Insecure     types.Bool   `tfsdk:"insecure"`
	Username     types.String `tfsdk:"username"`
	Password     types.String `tfsdk:"password"`
	Scheme       types.String `tfsdk:"scheme"`
	PollInterval types.Int64  `tfsdk:"poll_interval"`
}
