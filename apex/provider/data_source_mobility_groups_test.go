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

func TestAccMobilityGroupsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + mobilityCollectionConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of mobility_groups returned

					// Verify the first host to ensure all attributes are set
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_groups.test", "mobility_groups.0.name", "Create"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_groups.test", "mobility_groups.0.description", "Test"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_groups.test", "mobility_groups.0.system_id", "POWERFLEX-ELMVXRTEST0004"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_groups.test", "mobility_groups.0.system_type", "POWERFLEX"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_groups.test", "mobility_groups.0.creation_timestamp", "0001-01-01 00:00:00 +0000 UTC"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_groups.test", "mobility_groups.0.members.#", "1"),

					resource.TestCheckResourceAttr("data.apex_navigator_mobility_groups.test", "mobility_groups.0.members.0.id", "POWERFLEX-ELMVXRTEST0004__VOLUME__9e5a801700000000"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_groups.test", "mobility_groups.0.members.0.name", "Name"),
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_groups.test", "mobility_groups.0.members.0.size", "50"),
					// Verify placeholder id attribute
					resource.TestCheckResourceAttr("data.apex_navigator_mobility_groups.test", "id", "placeholder"),
				),
			},
			// Error reading mobility groups
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetMobilityGroupCollection).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + mobilityCollectionConfig,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Mobility Groups*.`),
			},
			// Filter testing single mobility group
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
				},
				Config: ProviderConfig + mobilityFilterSingleConfig,
			},
			// Filter testing multiple mobility group
			{
				Config: ProviderConfig + mobilityFilterMultipleConfig,
			},
			// Error getting mobility group with invalid filter
			{
				Config:      ProviderConfig + mobilityilterInvalidConfig,
				ExpectError: regexp.MustCompile(`.*one more more of the ids set in the filter is invalid*.`),
			},
		},
	})
}

var mobilityCollectionConfig = `data "apex_navigator_mobility_groups" "test" {}`
var mobilityFilterSingleConfig = `
 data "apex_navigator_mobility_groups" "example" {
	    filter {
	     ids = ["` + mobilityID1 + `"] 
	   }
	}
	
	output "instance_clone" {
	   value = data.apex_navigator_mobility_groups.example
	}
`

var mobilityFilterMultipleConfig = `
 data "apex_navigator_mobility_groups" "example" {
	     filter {
	     ids = ["` + mobilityID1 + `", "` + mobilityID2 + `"] 
	   }
	}
	
	output "instance_clone" {
	   value = data.apex_navigator_mobility_groups.example
	}
`

var mobilityilterInvalidConfig = `
 data "apex_navigator_mobility_groups" "example" {
	     filter {
	     ids = ["invalid-id"] 
	   }
	}
	
	output "instance_clone" {
	   value = data.apex_navigator_mobility_groups.example
	}
`
