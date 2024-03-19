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
				Config: ProviderConfig + storageProductsConfig,
				Check: resource.ComposeAggregateTestCheckFunc(

					// Verify the first storage product to ensure all attributes are set
					resource.TestCheckResourceAttr("data.apex_navigator_storage_products.test", "storage_products.0.cloud_type", "AWS"),
					resource.TestCheckResourceAttr("data.apex_navigator_storage_products.test", "storage_products.0.description", "Description12"),
					resource.TestCheckResourceAttr("data.apex_navigator_storage_products.test", "storage_products.0.latest_version", "LatestVersion"),
					resource.TestCheckResourceAttr("data.apex_navigator_storage_products.test", "storage_products.0.name", "NameStorage"),
					resource.TestCheckResourceAttr("data.apex_navigator_storage_products.test", "storage_products.0.storage_type", "BLOCK"),
					resource.TestCheckResourceAttr("data.apex_navigator_storage_products.test", "storage_products.0.system_type", "POWERFLEX"),
					resource.TestCheckResourceAttr("data.apex_navigator_storage_products.test", "storage_products.0.support_map.#", "1"),
					resource.TestCheckResourceAttr("data.apex_navigator_storage_products.test", "storage_products.0.support_map.0.id", "1"),
					resource.TestCheckResourceAttr("data.apex_navigator_storage_products.test", "storage_products.0.support_map.0.supported_actions.#", "0"),
					resource.TestCheckResourceAttr("data.apex_navigator_storage_products.test", "storage_products.0.support_map.0.supported_evaluation_period", "32"),
					resource.TestCheckResourceAttr("data.apex_navigator_storage_products.test", "storage_products.0.support_map.0.version", "VERSION"),

					// Verify placeholder id attribute
					resource.TestCheckResourceAttr("data.apex_navigator_storage_products.test", "id", "placeholder"),
				),
			},
			// Error Reading Storage Products
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetStorageProductsCollection).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + storageProductsConfig,
				ExpectError: regexp.MustCompile(`.*Unable Read to Storage Product Volume*.`),
			},
			// Filter testing single storage products
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
				},
				Config: ProviderConfig + storageProductsFilterSingleConfig,
			},
			// Filter testing multiple storage products
			{
				Config: ProviderConfig + storageProductsFilterMultipleConfig,
			},
			// Error getting filtered storage products
			{
				Config:      ProviderConfig + storageProductsFilterInvalidConfig,
				ExpectError: regexp.MustCompile(`.*one more more of the ids set in the filter is invalid*.`),
			},
		},
	})
}

var storageProductsConfig = `data "apex_navigator_storage_products" "test" {}`
var storageProductsFilterSingleConfig = `data "apex_navigator_storage_products" "test" {
	filter {
		ids = ["` + storageProductID1 + `"] 
	  }
}`
var storageProductsFilterMultipleConfig = `data "apex_navigator_storage_products" "test" {
	filter {
		ids = ["` + storageProductID1 + `", "` + storageProductID2 + `"] 
	  }
}`
var storageProductsFilterInvalidConfig = `data "apex_navigator_storage_products" "test" {
	filter {
		ids = ["invalid-id"] 
	}
}`
