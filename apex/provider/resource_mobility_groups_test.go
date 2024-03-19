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

func TestAccResourceMobilityGroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + moblilityGroupResourceConfig,
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
			{
				Config: ProviderConfig + moblilityGroupResourceUpdateConfig,
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
			{
				Config:      ProviderConfig + moblilityGroupResourceErrorUpdateConfig,
				ExpectError: regexp.MustCompile(`.*Error updating Mobility Group*.`),
			},
		},
	})
}

func TestAccResourceMobilityGroupError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create Mobility Group Error
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.CreateMobilityGroup).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + moblilityGroupResourceConfig,
				ExpectError: regexp.MustCompile(`.*Error creating Mobility Group*.`),
			},
			// Job failure error case
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.WaitForJobToComplete).Return(nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + moblilityGroupResourceConfig,
				ExpectError: regexp.MustCompile(`.*Error getting resourceID*.`),
			},
			// Attempt to get a mobility group after a job success error case
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.GetMobilityGroup).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + moblilityGroupResourceConfig,
				ExpectError: regexp.MustCompile(`.*Error retrieving created Mobility group*.`),
			},
			// Do a successful create to test update errors
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
				},
				Config: ProviderConfig + moblilityGroupResourceConfig,
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
			// Read Mobility Group Error
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.GetMobilityGroup).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + moblilityGroupResourceUpdateConfig,
				ExpectError: regexp.MustCompile(`.*Error Reading Apex Navigator mobility group*.`),
			},
			// Update Mobility Group Error
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.WaitForJobToComplete).Return(nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + moblilityGroupResourceUpdateConfig,
				ExpectError: regexp.MustCompile(`.*Error getting resourceID*.`),
			},
			// Update Mobility Group Job Error
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.UpdateMobilityGroup).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + moblilityGroupResourceConfig,
				ExpectError: regexp.MustCompile(`.*Error executing Update Mobility Group Job*.`),
			},
		},
	})
}

var moblilityGroupResourceConfig = `
resource "apex_navigator_mobility_groups" "example" {
	name        = "` + mobilityName + `" 
	system_id   = "` + systemID + `"
	system_type = "POWERFLEX"
	volume_id = [
	  "` + volumeID + `"
	]
  }
  
  output "examples_mobility_group" {
	value = apex_navigator_mobility_groups.example
  }
`

var moblilityGroupResourceUpdateConfig = `
resource "apex_navigator_mobility_groups" "example" {
	name        = "` + mobilityName + `_update" 
	system_id   = "` + systemID + `"
	system_type = "POWERFLEX"
	volume_id = [
	  "` + volumeID + `"
	]
  }
  
  output "examples_mobility_group" {
	value = apex_navigator_mobility_groups.example
  }
`

var moblilityGroupResourceErrorUpdateConfig = `
resource "apex_navigator_mobility_groups" "example" {
	name        = "` + mobilityName + `" 
	system_id   = "` + systemID + `"
	system_type = "POWERSCALE"
	volume_id = [
	  "` + volumeID + `"
	]
  }
  
  output "examples_mobility_group" {
	value = apex_navigator_mobility_groups.example
  }
`
