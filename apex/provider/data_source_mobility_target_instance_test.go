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

func TestAccMobilityTargetInstanceDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + configMoblilityTargetInstance,
				Check: resource.ComposeAggregateTestCheckFunc(

					// Verify the mobility target to ensure all attributes are set
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_target.test", "id", "POWERFLEX-ELMSIO0823QVTV__DATAMOBILITYGROUP__a90fcfaf-c61e-4b4d-8f89-65c6ef00dfd5"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_target.test", "name", "Get"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_target.test", "description", "Test"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_target.test", "system_id", "POWERFLEX-ELMVXRTEST0004"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_target.test", "system_type", "POWERFLEX"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_target.test", "source_mobility_group_id", "POWERFLEX-ELMSIO08200000__MOBILITYGROUP__3d07605a-0c68-4a86-9da2-dd9ccd1925af"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_target.test", "creation_timestamp", "0001-01-01 00:00:00 +0000 UTC"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_target.test", "image_timestamp", "0001-01-01 00:00:00 +0000 UTC"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_target.test", "bandwidth_limit", "512"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_target.test", "target_members.#", "1"),

					resource.TestCheckResourceAttr("data.apex_navigator_mobility_target.test", "target_members.0.id", "test"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_target.test", "target_members.0.name", "test"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_target.test", "target_members.0.parent_id", "POWERFLEX-ELMSIO0823QVTV__DATAMOBILITYGROUP__a90fcfaf-c61e-4b4d-8f89-65c6ef00dfd5"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_target.test", "target_members.0.size", "512"),
				),
			},
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetMobilityTarget).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + configMoblilityTargetInstance,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Mobility Target*.`),
			},
		},
	})
}

var configMoblilityTargetInstance = `data "apex_navigator_mobility_target" "test" {id="POWERFLEX-ELMSIO0823QVTV__DATAMOBILITYGROUP__a90fcfaf-c61e-4b4d-8f89-65c6ef00dfd5"}`
