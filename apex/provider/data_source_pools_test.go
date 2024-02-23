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

func TestAccPoolsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + sourcePoolsConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of pools returned
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.#", "1"),

					// Verify the first pool to ensure all attributes are set
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.free_size", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.id", "Pool_ID"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.issue_count", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.name", "Name"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.native_id", "Native_ID"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.subscribed_percent", "64.64"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.subscribed_size", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.system_id", "Pool_SID"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.system_model", "Sys_Model"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.system_name", "Sys_Name"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.system_type", "Pool_SystemType"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.time_to_full_prediction", "TTFP"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.total_size", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.type", "Type"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.used_percent", "64.64"),
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "pools.0.used_size", "64"),
					// Verify placeholder id attribute
					resource.TestCheckResourceAttr("data.apex_navigator_pools.test", "id", "placeholder"),
				),
			},
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetSourcePoolsCollection).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + sourcePoolsConfig,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Pools*.`),
			},
		},
	})
}

var sourcePoolsConfig = `data "apex_navigator_pools" "test" {}`
