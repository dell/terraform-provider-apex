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

func TestAccResourceMobilityGroupsCopy(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Create Mobility Group Copy
				Config: ProviderConfig + mobilityGroupsCopyResourceConfig,
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func TestAccResourceMobilityGroupsCopyError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Error in creating Mobility Group Copy
				PreConfig: func() {
					FunctionMocker = Mock(helper.CopyMobilityGroups).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + mobilityGroupsCopyResourceConfig,
				ExpectError: regexp.MustCompile(`.*Error creating Mobility Group Copy*.`),
			},
			{
				// Error while waiting for job to complete
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.WaitForJobToComplete).Return(nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + mobilityGroupsCopyResourceConfig,
				ExpectError: regexp.MustCompile(`.*Error getting resourceID*.`),
			},
			{
				// Error getting job status
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.GetJobStatus).Return(nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + mobilityGroupsCopyResourceConfig,
				ExpectError: regexp.MustCompile(`.*Error getting job*.`),
			},
		},
	})
}

var mobilityGroupsCopyResourceConfig = `
resource "apex_navigator_block_mobility_groups_copy" "example" {
	mobility_source_id = "` + mobilitySourceID + `"
	mobility_target_id = [
		"` + mobilityTargetID + `"
	]
  }
  
  output "examples_mobility_groups_copy" {
	value = apex_navigator_block_mobility_groups_copy.example
  }
`
