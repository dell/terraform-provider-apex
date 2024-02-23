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
	"fmt"
	"regexp"
	"testing"

	. "github.com/bytedance/mockey"
	"github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccClonesInstanceDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + cloneConfigInstance,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify the first host to ensure all attributes are set
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "id", "read_Id"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "image_timestamp", "read_Image Timestamp"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "mobility_target_id", "read_Mobility Target Id"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "name", "read_Name"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "refresh_timestamp", "read_Refresh Timestamp"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "creation_timestamp", "read_Creation Timestamp"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "description", "read_Description"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "clone_volumes.#", "1"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "clone_volumes.0.id", "read_VId"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "clone_volumes.0.name", "read_VName"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "clone_volumes.0.parent_id", "read_VParent Id"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "clone_volumes.0.size", "read_VSize"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "host_mappings.#", "1"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "host_mappings.0.host_id", "read_HId"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "host_mappings.0.host_initiator_protocol", "NVMe"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "host_mappings.0.host_ip", "read_HIp"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "host_mappings.0.host_name", "read_HName"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "host_mappings.0.id", "read_MapId"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "host_mappings.0.nqn", "read_Nqn"),
				),
			},
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetCloneInstance).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + cloneConfigInstance,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Clones*.`),
			},
		},
	})
}

var cloneConfigInstance = `data "apex_navigator_clone" "test" {id = "read_Id"}`
