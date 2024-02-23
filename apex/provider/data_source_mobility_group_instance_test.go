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

func TestMobilityGroupInstanceDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + moblilityGroupConfigInstance,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of storage_systems returned
					// resource.TestCheckResourceAttr("data.apex_navigator_test", "#", "1"),

					// // Verify the first host to ensure all attributes are set
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "id", "POWERFLEX-ELMSIO0823QVTV__DATAMOBILITYGROUP__fcdecfaf-c61e-4b4d-8f89-65c6ef00d0000"),

					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "id", "POWERFLEX-ELMSIO0823QVTV__DATAMOBILITYGROUP__fcdecfaf-c61e-4b4d-8f89-65c6ef00d0000"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "name", "Create"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "description", "Test"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "system_id", "POWERFLEX-ELMVXRTEST0004"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "system_type", "POWERFLEX"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "creation_timestamp", "0001-01-01 00:00:00 +0000 UTC"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "members.#", "1"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "members.0.id", "POWERFLEX-ELMVXRTEST0004__VOLUME__9e5a801700000000"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "members.0.name", "Name"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "members.0.size", "50"),
				),
			},
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetMobilityGroup).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + moblilityGroupConfigInstance,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Mobility Groups*.`),
			},
		},
	})
}

var moblilityGroupConfigInstance = `data "apex_navigator_mobility_group" "test" {id = "POWERFLEX-ELMSIO0823QVTV__DATAMOBILITYGROUP__fcdecfaf-c61e-4b4d-8f89-65c6ef00d0000"}`
