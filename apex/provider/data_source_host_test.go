/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://mozilla.org/MPL/2.0/

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package provider

import (
	"fmt"
	"regexp"
	"testing"

	. "github.com/bytedance/mockey"
	"github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccHostsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + hostConfig,
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
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetHostCollection).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + hostConfig,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Hosts*.`),
			},
		},
	})
}

var hostConfig = `data "apex_navigator_hosts" "test" {}`