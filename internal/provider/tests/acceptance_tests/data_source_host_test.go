package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccHostsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: sourceConfig + providerConfig + `data "apex_navigator_hosts" "test" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of hosts returned
					resource.TestCheckResourceAttr("data.apex_navigator_hosts.test", "hosts.#", "1"),

					// Verify the first host to ensure all attributes are set
					resource.TestCheckResourceAttr("data.apex_navigator_hosts.test", "hosts.0.description", "Example host description"),
					resource.TestCheckResourceAttr("data.apex_navigator_hosts.test", "hosts.0.id", "host123"),
					resource.TestCheckResourceAttr("data.apex_navigator_hosts.test", "hosts.0.initiator_count", "3"),
					resource.TestCheckResourceAttr("data.apex_navigator_hosts.test", "hosts.0.initiator_protocol", "TCP"),
					resource.TestCheckResourceAttr("data.apex_navigator_hosts.test", "hosts.0.issue_count", "2"),
					resource.TestCheckResourceAttr("data.apex_navigator_hosts.test", "hosts.0.name", "Example Host"),
					resource.TestCheckResourceAttr("data.apex_navigator_hosts.test", "hosts.0.native_id", "native789"),
					resource.TestCheckResourceAttr("data.apex_navigator_hosts.test", "hosts.0.network_addresses", "192.168.0.100"),
					resource.TestCheckResourceAttr("data.apex_navigator_hosts.test", "hosts.0.operating_system", "Windows 10"),
					resource.TestCheckResourceAttr("data.apex_navigator_hosts.test", "hosts.0.system_id", "sys456"),
					resource.TestCheckResourceAttr("data.apex_navigator_hosts.test", "hosts.0.system_model", "Dell XPS 15"),
					resource.TestCheckResourceAttr("data.apex_navigator_hosts.test", "hosts.0.system_name", "Example System"),
					resource.TestCheckResourceAttr("data.apex_navigator_hosts.test", "hosts.0.system_type", "Desktop"),
					resource.TestCheckResourceAttr("data.apex_navigator_hosts.test", "hosts.0.total_size", "500"),
					resource.TestCheckResourceAttr("data.apex_navigator_hosts.test", "hosts.0.type", "Virtual"),
					// Verify placeholder id attribute
					resource.TestCheckResourceAttr("data.apex_navigator_hosts.test", "id", "placeholder"),
				),
			},
		},
	})
}
