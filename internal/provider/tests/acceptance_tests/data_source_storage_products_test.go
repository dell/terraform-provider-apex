package tests // nolint:dupl

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccStorageProductsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: sourceConfig + providerConfig + `data "apex_navigator_storage_products" "test" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of storage products returned
					resource.TestCheckResourceAttr("data.apex_navigator_storage_products.test", "storage_products.#", "1"),

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
		},
	})
}
