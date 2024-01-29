package tests

// import (
// 	"testing"

// 	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
// )

// func TestAccStorageProductOptionsDataSource(t *testing.T) {
// 	resource.Test(t, resource.TestCase{
// 		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Read testing
// 			{
// 				Config: sourceConfig + providerConfig + `data "apex_navigator_storage_product_options" "test" {}`,
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					// Verify number of storage product options returned
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.#", "1"),

// 					// Verify the first host to ensure all attributes are set
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.id", "MyId"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.system_type", "POWERFLEX"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.storage_product_version", "v5.5 0"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.cloud_type", "AWS"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.#", "1"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.id", "MyID"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.name", "MyName"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.tier_type", "BALANCED"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.description", "Tier Description"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.storage_options.#", "1"),

// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.storage_options.0.id", "TierOptionId"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.storage_options.0.key", "MyKey"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.storage_options.0.value.#", "0"),

// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.storage_options.0.value.0.boolean_option.#", "0"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.storage_options.0.value.0.enum_option.#", "0"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.storage_options.0.value.0.free_text_option.#", "0"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.storage_options.0.value.0.multiselect_freetext_option.#", "0"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.storage_options.0.value.0.range_option.#", "0"),

// 					//todo: fix later when oneOf error when implementing the server
// 					// resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.cloud_options.#", "1"),
// 					// resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.cloud_options.0.id", "MyId"),
// 					// resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.cloud_options.0.key", "MyKey"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.cloud_options.0.value.#", "0"),

// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.cloud_options.0.value.0.boolean_option.#", "0"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.cloud_options.0.value.0.enum_option.#", "0"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.cloud_options.0.value.0.free_text_option.#", "0"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.cloud_options.0.value.0.multiselect_freetext_option.#", "0"),
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "storage_product_options.0.supported_tier_info.0.cloud_options.0.value.0.range_option.#", "0"),

// 					// Verify placeholder id attribute
// 					resource.TestCheckResourceAttr("data.apex_navigator_storage_product_options.test", "id", "placeholder"),
// 				),
// 			},
// 		},
// 	})
// }
