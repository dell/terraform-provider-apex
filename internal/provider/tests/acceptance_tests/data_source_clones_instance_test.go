package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccClonesInstanceDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: sourceConfig + providerConfig + `data "apex_navigator_clone" "test" {id = "read_Id"}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify the first host to ensure all attributes are set
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "id", "read_Id"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "image_timestamp", "read_Image Timestamp"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "mobility_target_id", "read_Mobility Target Id"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "name", "read_Name"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "refresh_timestamp", "read_Refresh Timestamp"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "creation_timestamp", "read_Creation Timestamp"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "description", "read_Description"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "clone_volumes.#", "1"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "clone_volumes.0.id", "read_VId"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "clone_volumes.0.name", "read_VName"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "clone_volumes.0.parent_id", "read_VParent Id"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "clone_volumes.0.size", "read_VSize"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "host_mappings.#", "1"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "host_mappings.0.host_id", "read_HId"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "host_mappings.0.host_initiator_protocol", "NVMe"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "host_mappings.0.host_ip", "read_HIp"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "host_mappings.0.host_name", "read_HName"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "host_mappings.0.id", "read_MapId"),
					resource.TestCheckResourceAttr("data.apex_navigator_clone.test", "host_mappings.0.nqn", "read_Nqn"),
				),
			},
		},
	})
}
