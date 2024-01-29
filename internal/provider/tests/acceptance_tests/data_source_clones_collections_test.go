package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccClonesDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: sourceConfig + providerConfig + `data "apex_navigator_clones" "test" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of clones returned
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.#", "1"),

					// Verify the first host to ensure all attributes are set
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.id", "read_Id"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.image_timestamp", "read_Image Timestamp"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.mobility_target_id", "read_Mobility Target Id"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.name", "read_Name"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.refresh_timestamp", "read_Refresh Timestamp"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.creation_timestamp", "read_Creation Timestamp"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.description", "read_Description"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.clone_volumes.#", "1"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.clone_volumes.0.id", "read_VId"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.clone_volumes.0.name", "read_VName"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.clone_volumes.0.parent_id", "read_VParent Id"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.clone_volumes.0.size", "read_VSize"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.host_mappings.#", "1"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.host_mappings.0.host_id", "read_HId"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.host_mappings.0.host_initiator_protocol", "NVMe"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.host_mappings.0.host_ip", "read_HIp"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.host_mappings.0.host_name", "read_HName"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.host_mappings.0.id", "read_MapId"),
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "clones.0.host_mappings.0.nqn", "read_Nqn"),
					// Verify placeholder id attribute
					resource.TestCheckResourceAttr("data.apex_navigator_clones.test", "id", "placeholder"),
				),
			},
		},
	})
}
