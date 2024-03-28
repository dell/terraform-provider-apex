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

// TestAccDataSourceVolumes is a Go function that tests the AccDataSourceVolumes functionality.
func TestAccDataSourcePools(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + sourcePoolsConfig + poolOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "true"),
					resource.TestCheckOutput("fetched_single", "false"),
				),
			},
			// Error getting volume collection
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetSourcePoolsCollection).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + sourcePoolsConfig + poolOutputs,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Pools*.`),
			},
			// Filter testing single volume
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
				},
				Config: ProviderConfig + sourcePoolsFilterSingleConfig + poolOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "false"),
					resource.TestCheckOutput("fetched_single", "true"),
					resource.TestCheckOutput("fetched_two", "false"),
				),
			},
			// Filter testing multiple volumes
			{
				Config: ProviderConfig + sourcePoolsFilterMultipleConfig + poolOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "true"),
					resource.TestCheckOutput("fetched_single", "false"),
					resource.TestCheckOutput("fetched_two", "true"),
				),
			},
			// Error getting filtered volume collection
			{
				Config:      ProviderConfig + sourcePoolsFilterInvalidConfig + poolOutputs,
				ExpectError: regexp.MustCompile(`.*one more more of the ids set in the filter is invalid*.`),
			},
		},
	})
}

var sourcePoolsConfig = `data "apex_navigator_block_pools" "test" {}`

var sourcePoolsFilterSingleConfig = `data "apex_navigator_block_pools" "test" {
	filter {
		ids = ["` + sourcePoolsID1 + `"] 
	  }
}`
var sourcePoolsFilterMultipleConfig = `data "apex_navigator_block_pools" "test" {
	filter {
		ids = ["` + sourcePoolsID1 + `", "` + sourcePoolsID2 + `"] 
	  }
}`
var sourcePoolsFilterInvalidConfig = `data "apex_navigator_block_pools" "test" {
	filter {
		ids = ["invalid-id"] 
	  }
}`

var poolOutputs = `
output "fetched_many" {
	value = length(data.apex_navigator_block_pools.test.pools) > 1
}
  
output "fetched_any" {
	value = length(data.apex_navigator_block_pools.test.pools) != 0
}

output "fetched_single" {
	value = length(data.apex_navigator_block_pools.test.pools) == 1
}

output "fetched_two" {
	value = length(data.apex_navigator_block_pools.test.pools) == 2
}
`
