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

func TestAccVolumesDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + volumeCollectionConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of volumes returned
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.#", "1"),

					// Verify the first host to ensure all attributes are set
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.allocated_size", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.bandwidth", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.consistency_group_name", "volume_gname"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.data_reduction_percent", "64.64"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.data_reduction_ratio", "64.64"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.data_reduction_saved_size", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.id", "volume_id"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.io_limit_policy_name", "volume_pname"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.iops", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.is_compressed_or_deduped", "volume_dduped"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.is_thin_enabled", "true"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.issue_count", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.latency", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.logical_size", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.name", "volume_name"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.native_id", "volume_nid"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.pool_id", "volume_poolid"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.pool_name", "volume_poolname"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.pool_type", "volume_pooltype"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.snap_shot_count", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.snap_shot_policy", "volume_spolicy"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.snap_shot_size", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.storage_resource_id", "volume_rid"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.storage_resource_native_id", "volume_rnid"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.system_id", "volume_sid"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.system_model", "volume_smodel"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.system_name", "volume_sname"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.system_type", "volume_stype"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.total_size", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.type", "volume_type"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.used_size", "64"),
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "volumes.0.used_size_unique", "64"),

					// Verify placeholder id attribute
					resource.TestCheckResourceAttr("data.apex_navigator_volumes.test", "id", "placeholder"),
				),
			},
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetVolumesCollection).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + volumeCollectionConfig,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Volumes*.`),
			},
		},
	})
}

var volumeCollectionConfig = `data "apex_navigator_volumes" "test" {}`
