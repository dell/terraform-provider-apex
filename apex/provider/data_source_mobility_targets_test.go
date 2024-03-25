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
				Config: ProviderConfig + moblilityTargetConfig + mobilityTargetsOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "true"),
					resource.TestCheckOutput("fetched_single", "false"),
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
				Config: ProviderConfig + mobilityTargetSingleFilterConfig + mobilityTargetsOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "false"),
					resource.TestCheckOutput("fetched_single", "true"),
					resource.TestCheckOutput("fetched_two", "false"),
				),
			},
			// Filter test for multiple mobility targets
			{
				Config: ProviderConfig + mobilityTargetMultipleFilterConfig + mobilityTargetsOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "true"),
					resource.TestCheckOutput("fetched_single", "false"),
					resource.TestCheckOutput("fetched_two", "true"),
				),
			},
			// Error getting mobility target for invalid filter
			{
				Config:      ProviderConfig + mobilityTargetInvalidFilterConfig + mobilityTargetsOutputs,
				ExpectError: regexp.MustCompile(`.*one more more of the ids set in the filter is invalid*.`),
			},
		},
	})
}

var moblilityTargetConfig = `data "apex_navigator_mobility_targets" "test" {}`
var mobilityTargetsOutputs = `
output "fetched_many" {
	value = length(data.apex_navigator_mobility_targets.test.mobility_targets) > 1
}
  
output "fetched_any" {
	value = length(data.apex_navigator_mobility_targets.test.mobility_targets) != 0
}

output "fetched_single" {
	value = length(data.apex_navigator_mobility_targets.test.mobility_targets) == 1
}

output "fetched_two" {
	value = length(data.apex_navigator_mobility_targets.test.mobility_targets) == 2
}
`
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
 data "apex_navigator_mobility_targets" "test" {
	     filter {
	     ids = ["invalid-id"] 
	   }
	}
	
	output "output_mobility_target" {
		value = data.apex_navigator_mobility_targets.test
	 }
`
