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

func TestAccResourceClone(t *testing.T) {
	var resTerraformName = "apex_navigator_clones.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read test
			{
				Config: ProviderConfig + cloneResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resTerraformName, "name", "CloneTerraformName"),
					resource.TestCheckResourceAttr(resTerraformName, "description", "Atlas Test BFF - test clone description."),
				),
			},
			// Import testing
			{
				ResourceName:      resTerraformName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update testing
			{
				Config: ProviderConfig + cloneResourceUpdateConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resTerraformName, "name", "CloneTerraformNameUpdated"),
					resource.TestCheckResourceAttr(resTerraformName, "description", "Atlas Test BFF - test clone description."),
				),
			},
		},
	})
}

func TestAccResourceCloneUpdateError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + cloneResourceConfig,
			},
			{
				Config:      ProviderConfig + cloneResourceUpdateErrorConfig,
				ExpectError: regexp.MustCompile(".*Could not update Clones*"),
			},
			{
				Config:      ProviderConfig + cloneResourceUpdateError2Config,
				ExpectError: regexp.MustCompile(".*Could not update Clones*"),
			},
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.UpdateClone).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + cloneResourceUpdateConfig,
				ExpectError: regexp.MustCompile(".*Could not execute update clones*"),
			},
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}

					FunctionMocker = Mock(helper.WaitForJobToComplete).Return(nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + cloneResourceUpdateConfig,
				ExpectError: regexp.MustCompile(`.*Error getting resourceID*.`),
			},
		},
	})
}
func TestAccResourceCloneCreateError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.CreateClone).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + cloneResourceConfig,
				ExpectError: regexp.MustCompile(".*Could not create Clones*"),
			},
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}

					FunctionMocker = Mock(helper.GetCloneInstance).Return(nil, &http.Response{Body: io.NopCloser(strings.NewReader("")), StatusCode: 400}, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + cloneResourceConfig,
				ExpectError: regexp.MustCompile(`.*Could not retrieve created Clone*.`),
			},
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}

					FunctionMocker = Mock(helper.WaitForJobToComplete).Return(nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + cloneResourceConfig,
				ExpectError: regexp.MustCompile(`.*Error getting resourceID*.`),
			},
		},
	})
}
func TestAccResourceCloneReadError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{

				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
				},
				Config: ProviderConfig + cloneResourceConfig,
			},
			{
				// read fail
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}

					FunctionMocker = Mock(helper.GetCloneInstance).Return(nil, &http.Response{Body: io.NopCloser(strings.NewReader("")), StatusCode: 400}, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + cloneResourceConfig,
				ExpectError: regexp.MustCompile(`.*Could not read Clones*.`),
			},
		},
	})
}

var cloneResourceConfig = `
resource "apex_navigator_clones" "example" {
	name               = "CloneTerraformName"
	description        = "Atlas Test BFF - test clone description."
	mobility_target_id = "POWERFLEX-ELMSIOENG10016__DATAMOBILITYGROUP__f71a0ef1-0a00-4c72-89e7-30d5180b1a0d"
	host_mappings = [
	  {
		host_id = "POWERFLEX-ELMSIOENG10016__HOST__5d743dc200000000"
	  }
	]
  }
`
var cloneResourceUpdateConfig = `
resource "apex_navigator_clones" "example" {
	name               = "CloneTerraformNameUpdated"
	description        = "Atlas Test BFF - test clone description."
	mobility_target_id = "POWERFLEX-ELMSIOENG10016__DATAMOBILITYGROUP__f71a0ef1-0a00-4c72-89e7-30d5180b1a0d"
	host_mappings = [
	  {
		host_id = "POWERFLEX-ELMSIOENG10016__HOST__5d743dc200000000"
	  }
	]
  }
`
var cloneResourceUpdateErrorConfig = `
				resource "apex_navigator_clones" "example" {
					name               = "CloneTerraformNameUpdateError"
					description        = "Atlas Test BFF - test clone description."
					mobility_target_id = "error"
					host_mappings = [
					  {
						host_id = "POWERFLEX-ELMSIOENG10016__HOST__5d743dc200000000"
					  }
					]
				  }
				
`
var cloneResourceUpdateError2Config = `
resource "apex_navigator_clones" "example" {
	name               = "CloneTerraformNameUpdatedError"
	description        = "Atlas Test BFF - test clone description."
	mobility_target_id = "error"
	host_mappings = [
	]
  }
`
