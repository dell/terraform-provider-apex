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

func TestAccDataSourceStorageProducts(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + storageProductsConfig + storageProductsOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "true"),
					resource.TestCheckOutput("fetched_single", "false"),
				),
			},
			// Error Reading Storage Products
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetStorageProductsCollection).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + storageProductsConfig,
				ExpectError: regexp.MustCompile(`.*Unable Read to storage product*.`),
			},
			// Filter testing single storage products
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
				},
				Config: ProviderConfig + storageProductsFilterSingleConfig + storageProductsOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "false"),
					resource.TestCheckOutput("fetched_single", "true"),
				),
			},
			// Error getting filtered storage products
			{
				Config:      ProviderConfig + storageProductsFilterInvalidConfig + storageProductsOutputs,
				ExpectError: regexp.MustCompile(`.*Failed to filter storage product*.`),
			},
		},
	})
}

var storageProductsConfig = `data "apex_navigator_storage_products" "test" {}`

var storageProductsOutputs = `
output "fetched_many" {
	value = length(data.apex_navigator_storage_products.test.storage_products) > 1
}
  
output "fetched_any" {
	value = length(data.apex_navigator_storage_products.test.storage_products) != 0
}

output "fetched_single" {
	value = length(data.apex_navigator_storage_products.test.storage_products) == 1
}

`

var storageProductsFilterSingleConfig = `data "apex_navigator_storage_products" "test" {
	filter {
		system_type = "` + storageProduct1 + `" 
	  }
}`

var storageProductsFilterInvalidConfig = `data "apex_navigator_storage_products" "test" {
	filter {
		system_type = "invalid-id"
	}
}`
