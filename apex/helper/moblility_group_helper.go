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
)

// GetMobilityGroupCollection returns a list of all Mobility Groups
func GetMobilityGroupCollection(client *client.APIClient) (*client.MobilityGroupsCollection200Response, *http.Response, error) {
	return client.MobilityGroupsAPI.MobilityGroupsCollection(context.Background()).Limit(500).Execute()
}

// GetMobilityGroup returns a single Mobility Group
func GetMobilityGroup(client *client.APIClient, id string) (*client.MobilityGroup, *http.Response, error) {
	return client.MobilityGroupsAPI.MobilityGroupsInstance(context.Background(), id).Execute()
}
