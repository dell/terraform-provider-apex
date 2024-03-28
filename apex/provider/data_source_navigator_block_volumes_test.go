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

// TestAccDataSourceVolumes is a Go function that tests the functionality of the AccDataSourceVolumes function.
//
// It takes a testing.T parameter and does not return any value.
func TestAccDataSourceVolumes(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + volumeCollectionConfig + volOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "true"),
					resource.TestCheckOutput("fetched_single", "false"),
				),
			},
			// Error getting volume collection
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetVolumesCollection).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + volumeCollectionConfig + volOutputs,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Volumes*.`),
			},
			// Filter testing single volume
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
				},
				Config: ProviderConfig + volumeCollectionFilterSingleConfig + volOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "false"),
					resource.TestCheckOutput("fetched_single", "true"),
					resource.TestCheckOutput("fetched_two", "false"),
				),
			},
			// Filter testing multiple volumes
			{
				Config: ProviderConfig + volumeCollectionFilterMultipleConfig + volOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "true"),
					resource.TestCheckOutput("fetched_single", "false"),
					resource.TestCheckOutput("fetched_two", "true"),
				),
			},
			// Error getting filtered volume collection
			{
				Config:      ProviderConfig + volumeCollectionFilterInvalidConfig + volOutputs,
				ExpectError: regexp.MustCompile(`.*one more more of the ids set in the filter is invalid*.`),
			},
		},
	})
}

var volumeCollectionConfig = `data "apex_navigator_block_volumes" "test" {}`
var volumeCollectionFilterSingleConfig = `data "apex_navigator_block_volumes" "test" {
	filter {
		ids = ["` + volumeID1 + `"] 
	  }
}`
var volumeCollectionFilterMultipleConfig = `data "apex_navigator_block_volumes" "test" {
	filter {
		ids = ["` + volumeID1 + `", "` + volumeID2 + `"] 
	  }
}`
var volumeCollectionFilterInvalidConfig = `data "apex_navigator_block_volumes" "test" {
	filter {
		ids = ["invalid-id"] 
	  }
}`

var volOutputs = `
output "fetched_many" {
	value = length(data.apex_navigator_block_volumes.test.volumes) > 1
}
  
output "fetched_any" {
	value = length(data.apex_navigator_block_volumes.test.volumes) != 0
}

output "fetched_single" {
	value = length(data.apex_navigator_block_volumes.test.volumes) == 1
}

output "fetched_two" {
	value = length(data.apex_navigator_block_volumes.test.volumes) == 2
}
`
