package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestStorageSystemInstanceDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: sourceConfig + providerConfig + `data "apex_navigator_block_storage" "test" {id = "POWERFLEX-ELMSIO0523STQ3-Mock"}`,
				Check: resource.ComposeAggregateTestCheckFunc(

					// // Verify the first host to ensure all attributes are set
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "id", "POWERFLEX-ELMSIO0523STQ3-Mock"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "system_id", "POWERFLEX-ELMSIO0523STQ3-Mock"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "system_type", "POWERFLEX"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "bandwidth", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "capacity_impact", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "capacity_issue_count", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "compression_savings", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "configuration_impact", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "configuration_issue_count", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "configured_size", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "connectivity_status", "CONNECTED"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "license_type", "AWS-license-mock"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "license_expiration_date_timestamp", "0001-01-01 00:00:00 +0000 UTC"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "contract_coverage_type", "covered"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "contract_expiration_date_timestamp", "0001-01-01 00:00:00 +0000 UTC"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "data_protection_impact", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "data_protection_issue_count", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "display_identifier", "ELMSIO0523STQ3"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "free_percent", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "free_size", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "health_connectivity_status", "NORMAL"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "health_issue_count", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "health_score", "99"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "health_state", "GOOD"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "iops", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "ipv4_address", "10.164.99.99"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "ipv6_address", "10.164.99.99"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "last_contact_timestamp", "0001-01-01 00:00:00 +0000 UTC"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "latency", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "logical_size", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "model", "PowerFlex software"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "name", "Create Cloud"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "overall_efficiency", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "performance_impact", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "performance_issue_count", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "serial_number", "serial-9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "site_name", "Round Rock"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "snaps_savings", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "system_health_impact", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "system_health_issue_count", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "thin_savings", "9999"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "total_size", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "unconfigured_size", "1"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "used_percent", "64.64"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "used_size", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "vendor", "Dell"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "version", "1.00"),

					// for deployment type PUBLIC_CLOUD
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "deployment_details.system_public_cloud.deployment_type", "PUBLIC_CLOUD"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "deployment_details.system_public_cloud.cloud_type", "AWS"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "deployment_details.system_public_cloud.cloud_account", "CloudAccount"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "deployment_details.system_public_cloud.cloud_region", "CloudRegion"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "deployment_details.system_public_cloud.availability_zone_topology", "MULTIPLE_AVAILABILITY_ZONE"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "deployment_details.system_public_cloud.virtual_private_cloud", "vpc-0c75759a9738c7f6f"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "deployment_details.system_public_cloud.cloud_management_address", "CloudManagementAddress"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "deployment_details.system_public_cloud.minimum_iops", "1000"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "deployment_details.system_public_cloud.minimum_capacity", "2199023255552"),
					resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "deployment_details.system_public_cloud.tier_type", "TierType"),

					// for deployment type ONPREM
					// resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "deployment_details.system_on_prem.site_name", "SiteName"),
					// resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "deployment_details.system_on_prem.location", "Location"),
					// resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "deployment_details.system_on_prem.country", "Country"),
					// resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "deployment_details.system_on_prem.state", "State"),
					// resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "deployment_details.system_on_prem.city", "City"),
					// resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "deployment_details.system_on_prem.street_address1", "StreetAddress1"),
					// resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "deployment_details.system_on_prem.street_address2", "StreetAddress2"),
					// resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "deployment_details.system_on_prem.zip_code", "ZipCode"),

					// Verify placeholder id attribute
					// resource.TestCheckResourceAttr("data.apex_navigator_block_storage.test", "id", "placeholder"),
				),
			},
		},
	})
}
