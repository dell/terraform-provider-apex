package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLicenseDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: sourceConfig + providerConfig + `data "apex_navigator_license" "test" {id = "licenseID"}`,
				Check: resource.ComposeAggregateTestCheckFunc(

					// Verify the first license to ensure all attributes are set
					resource.TestCheckResourceAttr("data.apex_navigator_license.test", "capacity", "3600000"),
					resource.TestCheckResourceAttr("data.apex_navigator_license.test", "capacity_units", "BYTE"),
					resource.TestCheckResourceAttr("data.apex_navigator_license.test", "entitlement_id", "licenseEntitlementID"),
					resource.TestCheckResourceAttr("data.apex_navigator_license.test", "expiration_type", "EXT_TRIAL"),
					resource.TestCheckResourceAttr("data.apex_navigator_license.test", "id", "licenseID"),
					resource.TestCheckResourceAttr("data.apex_navigator_license.test", "name", "licenseName"),
					resource.TestCheckResourceAttr("data.apex_navigator_license.test", "system_id", "licenseSystemID"),
					resource.TestCheckResourceAttr("data.apex_navigator_license.test", "type", "EVALUATION"),
				),
			},
		},
	})
}
