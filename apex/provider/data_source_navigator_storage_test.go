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

func TestAccDataSourceStorage(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + configStorage + storageOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "true"),
					resource.TestCheckOutput("fetched_single", "false"),
				),
			},
			// Error getting storage
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetStorageCollection).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + configStorage,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Storage*.`),
			},
			// Filter testing single storage
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
				},
				Config: ProviderConfig + configFilteredSingleStorage + storageOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "false"),
					resource.TestCheckOutput("fetched_single", "true"),
					resource.TestCheckOutput("fetched_two", "false"),
				),
			},
			// Filter testing multiple storage
			{
				Config: ProviderConfig + configFilteredMultipleStorage + storageOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "true"),
					resource.TestCheckOutput("fetched_single", "false"),
					resource.TestCheckOutput("fetched_two", "true"),
				),
			},
			// Filter by system type
			{
				Config: ProviderConfig + configFilteredSystemType + storageOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "true"),
					resource.TestCheckOutput("fetched_single", "false"),
				),
			},
			// Error getting storage
			{
				Config:      ProviderConfig + configFilteredInvalidStorage,
				ExpectError: regexp.MustCompile(`.*one more more of the ids set in the filter is invalid*.`),
			},
		},
	})
}

var configStorage = `data "apex_navigator_storages" "test" {}`

var storageOutputs = `
output "fetched_many" {
	value = length(data.apex_navigator_storages.test.storages) > 1
}
  
output "fetched_any" {
	value = length(data.apex_navigator_storages.test.storages) != 0
}

output "fetched_single" {
	value = length(data.apex_navigator_storages.test.storages) == 1
}

output "fetched_two" {
	value = length(data.apex_navigator_storages.test.storages) == 2
}
`

var configFilteredSystemType = `
data "apex_navigator_storages" "test" {
	   filter {
			system_type = "POWERFLEX" 
	   }
}`

var configFilteredSingleStorage = `
data "apex_navigator_storages" "test" {
	   filter {
	     ids = ["` + blockStorageID1 + `"] 
	   }
}`

var configFilteredMultipleStorage = `
data "apex_navigator_storages" "test" {
	   filter {
	     ids = ["` + blockStorageID1 + `", "` + blockStorageID2 + `"] 
	   }
}`

var configFilteredInvalidStorage = `
data "apex_navigator_storages" "test" {
	   filter {
	     ids = ["invalid-id"] 
	   }
}`
