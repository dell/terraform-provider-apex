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

func TestAccDataSourceHosts(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + hostConfig + hostOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "true"),
					resource.TestCheckOutput("fetched_single", "false"),
				),
			},
			// Error getting host collection
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetHostCollection).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + hostConfig + hostOutputs,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Hosts*.`),
			},
			// Filter testing single host
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
				},
				Config: ProviderConfig + hostFilterSingleConfig + hostOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "false"),
					resource.TestCheckOutput("fetched_single", "true"),
					resource.TestCheckOutput("fetched_two", "false"),
				),
			},
			// Filter testing multiple hosts
			{
				Config: ProviderConfig + hostFilterMultipleConfig + hostOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "true"),
					resource.TestCheckOutput("fetched_single", "false"),
					resource.TestCheckOutput("fetched_two", "true"),
				),
			},
			// Error getting filtered host collection
			{
				Config:      ProviderConfig + hostFilterInvalidConfig + hostOutputs,
				ExpectError: regexp.MustCompile(`.*one more more of the ids set in the filter is invalid*.`),
			},
		},
	})
}

var hostConfig = `data "apex_navigator_block_hosts" "example" {}`

var hostOutputs = `
output "fetched_many" {
	value = length(data.apex_navigator_block_hosts.example.hosts) > 1
}
  
output "fetched_any" {
	value = length(data.apex_navigator_block_hosts.example.hosts) != 0
}

output "fetched_single" {
	value = length(data.apex_navigator_block_hosts.example.hosts) == 1
}

output "fetched_two" {
	value = length(data.apex_navigator_block_hosts.example.hosts) == 2
}
`

var hostFilterSingleConfig = `
 data "apex_navigator_block_hosts" "example" {
	    filter {
	     ids = ["` + hostID1 + `"] 
	   }
	}
`

var hostFilterMultipleConfig = `
 data "apex_navigator_block_hosts" "example" {
	     filter {
	     ids = ["` + hostID1 + `", "` + hostID2 + `"] 
	   }
	}
`

var hostFilterInvalidConfig = `
 data "apex_navigator_block_hosts" "example" {
	     filter {
	     ids = ["invalid-id"] 
	   }
	}
`
