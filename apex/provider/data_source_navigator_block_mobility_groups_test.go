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

func TestAccDataSourceMobilityGroups(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + mobilityCollectionConfig + mobilityGroupsOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "true"),
					resource.TestCheckOutput("fetched_single", "false"),
				),
			},
			// Error reading mobility groups
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetMobilityGroupCollection).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + mobilityCollectionConfig + mobilityGroupsOutputs,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Mobility Groups*.`),
			},
			// Filter testing single mobility group
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
				},
				Config: ProviderConfig + mobilityFilterSingleConfig + mobilityGroupsOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "false"),
					resource.TestCheckOutput("fetched_single", "true"),
					resource.TestCheckOutput("fetched_two", "false"),
				),
			},
			// Filter testing multiple mobility group
			{
				Config: ProviderConfig + mobilityFilterMultipleConfig + mobilityGroupsOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "true"),
					resource.TestCheckOutput("fetched_single", "false"),
					resource.TestCheckOutput("fetched_two", "true"),
				),
			},
			// Error getting mobility group with invalid filter
			{
				Config:      ProviderConfig + mobilityilterInvalidConfig + mobilityGroupsOutputs,
				ExpectError: regexp.MustCompile(`.*one more more of the ids set in the filter is invalid*.`),
			},
		},
	})
}

var mobilityCollectionConfig = `data "apex_navigator_block_mobility_groups" "example" {}`

var mobilityGroupsOutputs = `
output "fetched_many" {
	value = length(data.apex_navigator_block_mobility_groups.example.mobility_groups) > 1
}
  
output "fetched_any" {
	value = length(data.apex_navigator_block_mobility_groups.example.mobility_groups) != 0
}

output "fetched_single" {
	value = length(data.apex_navigator_block_mobility_groups.example.mobility_groups) == 1
}

output "fetched_two" {
	value = length(data.apex_navigator_block_mobility_groups.example.mobility_groups) == 2
}
`

var mobilityFilterSingleConfig = `
 data "apex_navigator_block_mobility_groups" "example" {
	    filter {
	     ids = ["` + mobilityID1 + `"] 
	   }
	}
	
	output "instance_clone" {
	   value = data.apex_navigator_block_mobility_groups.example
	}
`

var mobilityFilterMultipleConfig = `
 data "apex_navigator_block_mobility_groups" "example" {
	     filter {
	     ids = ["` + mobilityID1 + `", "` + mobilityID2 + `"] 
	   }
	}
	
	output "instance_clone" {
	   value = data.apex_navigator_block_mobility_groups.example
	}
`

var mobilityilterInvalidConfig = `
 data "apex_navigator_block_mobility_groups" "example" {
	     filter {
	     ids = ["invalid-id"] 
	   }
	}
	
	output "instance_clone" {
	   value = data.apex_navigator_block_mobility_groups.example
	}
`
