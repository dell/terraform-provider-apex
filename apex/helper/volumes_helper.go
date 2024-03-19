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

// GetVolumesCollection returns a list of all Volumes
func GetVolumesCollection(client *client.APIClient, filter string) (*client.VolumesCollection200Response, *http.Response, error) {
	return client.VolumesAPI.VolumesCollection(context.Background()).Filter(filter).Limit(500).Execute()
}
