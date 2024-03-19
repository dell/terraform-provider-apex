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

func TestAccDataSourceMobilityTargets(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + moblilityTargetConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of mobility_targets returned
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.#", "2"),

					// Verify the first host to ensure all attributes are set
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.id", "POWERFLEX-ELMSIO0823QVTV__DATAMOBILITYGROUP__a90fcfaf-c61e-4b4d-8f89-65c6ef00dfd5"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.name", "Get"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.description", "Test"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.system_id", "POWERFLEX-ELMVXRTEST0004"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.system_type", "POWERFLEX"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.source_mobility_group_id", "POWERFLEX-ELMSIO08200000__MOBILITYGROUP__3d07605a-0c68-4a86-9da2-dd9ccd1925af"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.creation_timestamp", "0001-01-01 00:00:00 +0000 UTC"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.image_timestamp", "0001-01-01 00:00:00 +0000 UTC"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.bandwidth_limit", "512"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.target_members.#", "1"),

					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.target_members.0.id", "test"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.target_members.0.name", "test"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.target_members.0.parent_id", "POWERFLEX-ELMSIO0823QVTV__DATAMOBILITYGROUP__a90fcfaf-c61e-4b4d-8f89-65c6ef00dfd5"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.target_members.0.size", "512"),

					// Verify placeholder id attribute
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "id", "placeholder"),
				),
			},
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetMobilityTargetCollection).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + moblilityTargetConfig,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Mobility Target*.`),
			},
			// Filter test for single mobility target
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
				},
				Config: ProviderConfig + mobilityTargetSingleFilterConfig,
			},
			// Filter test for multiple mobility targets
			{
				Config: ProviderConfig + mobilityTargetMultipleFilterConfig,
			},
			// Error getting mobility target for invalid filter
			{
				Config:      ProviderConfig + mobilityTargetInvalidFilterConfig,
				ExpectError: regexp.MustCompile(`.*one more more of the ids set in the filter is invalid*.`),
			},
		},
	})
}

var moblilityTargetConfig = `data "apex_navigator_mobility_targets" "test" {}`
var mobilityTargetSingleFilterConfig = `
 data "apex_navigator_mobility_targets" "test" {
	    filter {
	     ids = ["` + mobilityTargetID1 + `"] 
	   }
	}
	
	output "output_mobility_target" {
	   value = data.apex_navigator_mobility_targets.test
	}
`

var mobilityTargetMultipleFilterConfig = `
 data "apex_navigator_mobility_targets" "test" {
	     filter {
	     ids = ["` + mobilityTargetID1 + `", "` + mobilityTargetID2 + `"] 
	   }
	}
	
	output "output_mobility_target" {
	   value = data.apex_navigator_mobility_targets.test
	}
`

var mobilityTargetInvalidFilterConfig = `
 data "apex_navigator_mobility_targets" "example" {
	     filter {
	     ids = ["invalid-id"] 
	   }
	}
	
	output "output_mobility_target" {
		value = data.apex_navigator_mobility_targets.example
	 }
`
