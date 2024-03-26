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

package helper

import (
	"context"
	"net/http"

	client "dell/apex-client"

	"github.com/dell/terraform-provider-apex/apex/models"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// GetHostCollection gets list of hosts.
func GetHostCollection(client *client.APIClient, filter string) (*client.HostsCollection200Response, *http.Response, error) {
	return client.HostsAPI.HostsCollection(context.Background()).Filter(filter).Limit(500).Execute()
}

// MapHostModel maps response body to model
func MapHostModel(hosts *client.HostsCollection200Response, state models.HostsDataSourceModel) models.HostsDataSourceModel {
	// Map response body to model
	for _, host := range hosts.Results {
		hostState := models.HostsModel{
			ID:                types.StringValue(host.Id),
			SystemID:          types.StringPointerValue(host.SystemId),
			SystemType:        types.StringPointerValue(host.SystemType),
			Description:       types.StringPointerValue(host.Description),
			InitiatorCount:    types.Int64PointerValue(host.InitiatorCount),
			InitiatorProtocol: types.StringPointerValue(host.InitiatorProtocol),
			IssueCount:        types.Int64PointerValue(host.IssueCount),
			Name:              types.StringPointerValue(host.Name),
			NativeID:          types.StringPointerValue(host.NativeId),
			NetworkAddresses:  types.StringPointerValue(host.NetworkAddresses),
			Type:              types.StringPointerValue(host.Type),
			OperatingSystem:   types.StringPointerValue(host.OperatingSystem),
			SystemModel:       types.StringPointerValue(host.SystemModel),
			SystemName:        types.StringPointerValue(host.SystemName),
			TotalSize:         types.Int64PointerValue(host.TotalSize),
		}

		state.Hosts = append(state.Hosts, hostState)
	}
	return state
}
