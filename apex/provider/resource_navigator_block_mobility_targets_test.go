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

// Acceptance Tests

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"testing"

	. "github.com/bytedance/mockey"

	"github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceMobilityTargets(t *testing.T) {
	var resTerraformName = "apex_navigator_block_mobility_targets.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read test
			{
				Config: ProviderConfig + mobilityTargetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resTerraformName, "name", "create-mobility-target"),
					resource.TestCheckResourceAttr(resTerraformName, "system_type", "POWERFLEX"),
				),
			},
			//Import testing
			{
				ResourceName:  resTerraformName,
				ImportStateId: "{\"id\":\"" + mobilityTargetID + "\",\"username\":\"" + powerflexUser + "\",\"password\":\"" + powerflexPass + "\",\"host\":\"" + "" + "\",\"insecure\":true}",
				ImportState:   true,
			},
			//Import Error testing
			{
				ResourceName:  resTerraformName,
				ImportStateId: "{\"username\":\"" + powerflexUser + "\",\"password\":\"" + powerflexPass + "\",\"host\":\"" + "" + "\",\"insecure\":true}",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`.*Could not retrieve Mobility Target during import*.`),
			},
			//Update testing
			{
				Config: ProviderConfig + mobilityTargetResourceUpdateConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resTerraformName, "name", "update-mobility-target"),
					resource.TestCheckResourceAttr(resTerraformName, "system_type", "POWERFLEX"),
				),
			},
		},
	})
}

func TestAccResourceMobilityTargetsUpdateError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + mobilityTargetResourceConfig,
			},
			// Activate Powerflex Error
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.ActivateSystemClientSystem).Return(fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + mobilityTargetResourceUpdateConfig,
				ExpectError: regexp.MustCompile(`.*Error activating Powerflex System*.`),
			},
			{
				// Read failure
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}

					FunctionMocker = Mock(helper.GetMobilityTarget).Return(nil, &http.Response{Body: io.NopCloser(strings.NewReader("")), StatusCode: 400}, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + mobilityTargetResourceUpdateConfig,
				ExpectError: regexp.MustCompile(`.*Could not read Mobility target*.`),
			},
			{
				// Update failure
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.UpdateMobilityTarget).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + mobilityTargetResourceUpdateConfig,
				ExpectError: regexp.MustCompile(".*Error executing Update Mobility Target Job*"),
			},

			{
				// Error while waiting for job after update
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}

					FunctionMocker = Mock(helper.WaitForJobToComplete).Return(nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + mobilityTargetResourceUpdateConfig,
				ExpectError: regexp.MustCompile(`.*Error getting resourceID*.`),
			},
		},
	})
}

func TestAccResourceMobilityTargetsCreateError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Create failure
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.CreateMobilityTarget).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + mobilityTargetResourceConfig,
				ExpectError: regexp.MustCompile(".*Could not create Mobility Target*"),
			},
			// Activate Powerflex Error
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.ActivateSystemClientSystem).Return(fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + mobilityTargetResourceConfig,
				ExpectError: regexp.MustCompile(`.*Error activating Powerflex System*.`),
			},
			{
				// Read failure after create
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}

					FunctionMocker = Mock(helper.GetMobilityTarget).Return(nil, &http.Response{Body: io.NopCloser(strings.NewReader("")), StatusCode: 400}, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + mobilityTargetResourceConfig,
				ExpectError: regexp.MustCompile(`.*Could not retrieve Mobility group*.`),
			},
			{
				// Error while waiting for job after create
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}

					FunctionMocker = Mock(helper.WaitForJobToComplete).Return(nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + mobilityTargetResourceConfig,
				ExpectError: regexp.MustCompile(`.*Error getting resourceID*.`),
			},
		},
	})
}

var mobilityTargetResourceConfig = `
resource "apex_navigator_block_mobility_targets" "example" {
	name                     = "create-mobility-target"
	source_mobility_group_id = "` + sourceMobilityTargetGroupID + `"
	system_id                = "` + mobilityTargetSystemID + `"
	system_type              = "POWERFLEX"
	target_system_options    = "` + mobilityTargetOptions + `"
	powerflex {
		username = "` + powerflexUser + `"
		password = "` + powerflexPass + `"
		scheme   = "` + powerflexScheme + `"   
	}
  }
`
var mobilityTargetResourceUpdateConfig = `
resource "apex_navigator_block_mobility_targets" "example" {
	name                     = "update-mobility-target"
	source_mobility_group_id = "` + sourceMobilityTargetGroupID + `"
	system_id                = "` + mobilityTargetSystemID + `"
	system_type              = "POWERFLEX"
	target_system_options    = "` + mobilityTargetOptions + `"
	powerflex {
		username = "` + powerflexUser + `"
		password = "` + powerflexPass + `"
		scheme   = "` + powerflexScheme + `"   
	}
  }
`
