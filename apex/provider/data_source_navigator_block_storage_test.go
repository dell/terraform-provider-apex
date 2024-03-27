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

func TestAccDataSourceBlockStorage(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + configBlockStorage + blockStorageOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "true"),
					resource.TestCheckOutput("fetched_single", "false"),
				),
			},
			// Error getting block storage
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetBlockStorageCollection).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + configBlockStorage,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Block Storage*.`),
			},
			// Filter testing single block storage
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
				},
				Config: ProviderConfig + configFilteredBlockSingleStorage + blockStorageOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "false"),
					resource.TestCheckOutput("fetched_single", "true"),
					resource.TestCheckOutput("fetched_two", "false"),
				),
			},
			// Filter testing multiple block storage
			{
				Config: ProviderConfig + configFilteredBlockMultipleStorage + blockStorageOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "true"),
					resource.TestCheckOutput("fetched_single", "false"),
					resource.TestCheckOutput("fetched_two", "true"),
				),
			},
			// Error getting block storage
			{
				Config:      ProviderConfig + configFilteredInvalidStorage,
				ExpectError: regexp.MustCompile(`.*one more more of the ids set in the filter is invalid*.`),
			},
		},
	})
}

var configBlockStorage = `data "apex_navigator_block_storages" "test" {}`

var blockStorageOutputs = `
output "fetched_many" {
	value = length(data.apex_navigator_block_storages.test.block_storages) > 1
}
  
output "fetched_any" {
	value = length(data.apex_navigator_block_storages.test.block_storages) != 0
}

output "fetched_single" {
	value = length(data.apex_navigator_block_storages.test.block_storages) == 1
}

output "fetched_two" {
	value = length(data.apex_navigator_block_storages.test.block_storages) == 2
}
`

var configFilteredBlockSingleStorage = `
data "apex_navigator_block_storages" "test" {
	   filter {
	     ids = ["` + blockStorageID1 + `"] 
	   }
}`

var configFilteredBlockMultipleStorage = `
data "apex_navigator_block_storages" "test" {
	   filter {
	     ids = ["` + blockStorageID1 + `", "` + blockStorageID2 + `"] 
	   }
}`

var configFilteredInvalidStorage = `
data "apex_navigator_block_storages" "test" {
	   filter {
	     ids = ["invalid-id"] 
	   }
}`
