package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccPoolsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: sourceConfig + providerConfig + `data "apex_navigator_pools" "test" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of pools returned
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.#", "1"),

					// Verify the first pool to ensure all attributes are set
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.free_size", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.id", "Pool_ID"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.issue_count", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.name", "Name"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.native_id", "Native_ID"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.subscribed_percent", "64.64"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.subscribed_size", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.system_id", "Pool_SID"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.system_model", "Sys_Model"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.system_name", "Sys_Name"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.system_type", "Pool_SystemType"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.time_to_full_prediction", "TTFP"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.total_size", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.type", "Type"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.used_percent", "64.64"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.used_size", "64"),
					// Verify placeholder id attribute
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "id", "placeholder"),
				),
			},
		},
	})
}
