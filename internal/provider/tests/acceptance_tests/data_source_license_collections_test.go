package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLicensesDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: sourceConfig + providerConfig + `data "apex_navigator_licenses" "test" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of licenses returned
					resource.TestCheckResourceAttr("data.apex_navigator_licenses.test", "licenses.#", "1"),

					// Verify the first license to ensure all attributes are set
					resource.TestCheckResourceAttr("data.apex_navigator_licenses.test", "licenses.0.capacity", "3600000"),
					resource.TestCheckResourceAttr("data.apex_navigator_licenses.test", "licenses.0.capacity_units", "BYTE"),
					resource.TestCheckResourceAttr("data.apex_navigator_licenses.test", "licenses.0.entitlement_id", "licenseEntitlementID"),
					resource.TestCheckResourceAttr("data.apex_navigator_licenses.test", "licenses.0.expiration_type", "EXT_TRIAL"),
					resource.TestCheckResourceAttr("data.apex_navigator_licenses.test", "licenses.0.id", "licenseID"),
					resource.TestCheckResourceAttr("data.apex_navigator_licenses.test", "licenses.0.name", "licenseName"),
					resource.TestCheckResourceAttr("data.apex_navigator_licenses.test", "licenses.0.system_id", "licenseSystemID"),
					resource.TestCheckResourceAttr("data.apex_navigator_licenses.test", "licenses.0.type", "EVALUATION"),
					// Verify placeholder id attribute
					resource.TestCheckResourceAttr("data.apex_navigator_licenses.test", "id", "placeholder"),
				),
			},
		},
	})
}
