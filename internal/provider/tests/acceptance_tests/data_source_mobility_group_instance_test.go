package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestMobilityGroupInstanceDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: sourceConfig + providerConfig + `data "apex_navigator_mobility_group" "test" {id = "POWERFLEX-ELMSIO0823QVTV__DATAMOBILITYGROUP__fcdecfaf-c61e-4b4d-8f89-65c6ef00d0000"}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of storage_systems returned
					// resource.TestCheckResourceAttr("data.apex_navigator_test", "#", "1"),

					// // Verify the first host to ensure all attributes are set
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "id", "POWERFLEX-ELMSIO0823QVTV__DATAMOBILITYGROUP__fcdecfaf-c61e-4b4d-8f89-65c6ef00d0000"),

					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "id", "POWERFLEX-ELMSIO0823QVTV__DATAMOBILITYGROUP__fcdecfaf-c61e-4b4d-8f89-65c6ef00d0000"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "name", "Create"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "description", "Test"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "system_id", "POWERFLEX-ELMVXRTEST0004"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "system_type", "POWERFLEX"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "creation_timestamp", "0001-01-01 00:00:00 +0000 UTC"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "members.#", "1"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "members.0.id", "POWERFLEX-ELMVXRTEST0004__VOLUME__9e5a801700000000"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "members.0.name", "Name"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_group.test", "members.0.size", "50"),
				),
			},
		},
	})
}
