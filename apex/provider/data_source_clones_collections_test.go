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

func TestAccClonesDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + cloneConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of clones returned
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.#", "1"),

					// Verify the first host to ensure all attributes are set
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.id", "read_Id"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.image_timestamp", "read_Image Timestamp"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.mobility_target_id", "read_Mobility Target Id"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.name", "read_Name"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.refresh_timestamp", "read_Refresh Timestamp"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.creation_timestamp", "read_Creation Timestamp"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.description", "read_Description"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.clone_volumes.#", "1"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.clone_volumes.0.id", "read_VId"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.clone_volumes.0.name", "read_VName"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.clone_volumes.0.parent_id", "read_VParent Id"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.clone_volumes.0.size", "read_VSize"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.host_mappings.#", "1"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.host_mappings.0.host_id", "read_HId"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.host_mappings.0.host_initiator_protocol", "NVMe"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.host_mappings.0.host_ip", "read_HIp"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.host_mappings.0.host_name", "read_HName"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.host_mappings.0.id", "read_MapId"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.host_mappings.0.nqn", "read_Nqn"),
					// Verify placeholder id attribute
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "id", "placeholder"),
				),
			},
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetCloneCollection).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + cloneConfig,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Clones*.`),
			},
		},
	})
}

var cloneConfig = `data "apex_navigator_clones" "test" {}`
