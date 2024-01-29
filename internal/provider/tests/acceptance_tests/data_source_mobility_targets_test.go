package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccMobilityTargetsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: sourceConfig + providerConfig + `data "apex_navigator_mobility_targets" "test" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of mobility_targets returned
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.#", "2"),

					// Verify the first host to ensure all attributes are set
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.id", "POWERFLEX-ELMSIO0823QVTV__DATAMOBILITYGROUP__a90fcfaf-c61e-4b4d-8f89-65c6ef00dfd5"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.name", "Get"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.description", "Test"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.system_id", "POWERFLEX-ELMVXRTEST0004"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.system_type", "POWERFLEX"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.source_mobility_group_id", "POWERFLEX-ELMSIO08200000__MOBILITYGROUP__3d07605a-0c68-4a86-9da2-dd9ccd1925af"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.creation_timestamp", "0001-01-01 00:00:00 +0000 UTC"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.image_timestamp", "0001-01-01 00:00:00 +0000 UTC"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.bandwidth_limit", "512"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.target_members.#", "1"),

					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.target_members.0.id", "test"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.target_members.0.name", "test"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.target_members.0.parent_id", "POWERFLEX-ELMSIO0823QVTV__DATAMOBILITYGROUP__a90fcfaf-c61e-4b4d-8f89-65c6ef00dfd5"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "mobility_targets.0.target_members.0.size", "512"),

					// Verify placeholder id attribute
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_targets.test", "id", "placeholder"),
				),
			},
		},
	})
}
