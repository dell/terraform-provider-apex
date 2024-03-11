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

	client "github.com/dell/terraform-provider-apex/client/apexclient/client"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// GetCloneCollection is a collection of clones
func GetCloneCollection(client *client.APIClient) (*client.ClonesCollection200Response, *http.Response, error) {
	return client.ClonesAPI.ClonesCollection(context.Background()).Limit(500).Execute()
}

// GetCloneInstance is a clone
func GetCloneInstance(client *client.APIClient, id string) (*client.Clone, *http.Response, error) {
	return client.ClonesAPI.ClonesInstance(context.Background(), id).Execute()
}

// CreateClone sends the create clone request
func CreateClone(request client.ApiClonesCreateRequest, input client.CloneCreateInput) (*client.Job, *http.Response, error) {
	request = request.CloneCreateInput(input)
	return request.Async(true).Execute()
}

// UpdateClone sends the update clone request
func UpdateClone(request client.ApiClonesModifyRequest, input client.UpdateCloneInput) (*client.Job, *http.Response, error) {
	request = request.UpdateCloneInput(input)
	return request.Async(true).Execute()
}

// MapClones maps clones to a host
func MapClones(ctx context.Context, mapReq client.ApiClonesMapRequest, cloneID string, hosts []basetypes.StringValue) (*client.Job, *http.Response, error) {
	// Create Clones POST request
	hostIds := make([]string, 0, len(hosts))
	for _, mapping := range hosts {
		hostIds = append(hostIds, mapping.ValueString())
	}
	mapInput := *client.NewMapInput(hostIds)
	mapReq = mapReq.MapInput(mapInput)
	return mapReq.Async(true).Execute()
}

// RefreshClone refreshes a clone
func RefreshClone(refreshReq client.ApiClonesRefreshRequest) (*client.Job, *http.Response, error) {
	// Executing refresh request request
	return refreshReq.Async(true).Execute()
}

// UnmapClones unmaps clones from a host
func UnmapClones(unmapReq client.ApiClonesUnmapRequest, hosts []basetypes.StringValue) (*client.Job, *http.Response, error) {
	hostIds := make([]string, 0, len(hosts))
	for _, mapping := range hosts {
		hostIds = append(hostIds, mapping.ValueString())
	}
	unmapInput := *client.NewUnmapInput(hostIds)
	unmapReq = unmapReq.UnmapInput(unmapInput)
	return unmapReq.Async(true).Execute()
}
