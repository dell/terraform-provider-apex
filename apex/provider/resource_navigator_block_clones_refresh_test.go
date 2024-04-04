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

func TestAccResourceClonesRefresh(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + clonesRefreshResourceConfig,
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
			// Updates should pass since it is just a passthrough
			{
				Config: ProviderConfig + clonesRefreshResourceConfig,
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func TestAccResourceClonesRefreshError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.RefreshClone).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + clonesRefreshResourceConfig,
				ExpectError: regexp.MustCompile(`.*Error creating Clones Refresh request*.`),
			},
			// Activate Powerflex Error
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.ActivateSystemPowerflexSystem).Return(fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + clonesRefreshResourceConfig,
				ExpectError: regexp.MustCompile(`.*Error activating Powerflex System*.`),
			},
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.WaitForJobToComplete).Return(nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + clonesRefreshResourceConfig,
				ExpectError: regexp.MustCompile(`.*Error getting resourceID*.`),
			},
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.GetJobStatus).Return(nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + clonesRefreshResourceConfig,
				ExpectError: regexp.MustCompile(`.*Error getting job*.`),
			},
		},
	})
}

var clonesRefreshResourceConfig = `
resource "apex_navigator_block_clones_refresh" "example" {
	clone_id = "` + cloneRefreshID + `"
	system_id = "` + systemID + `"
	powerflex {
		username = "` + powerflexUser + `"
		password = "` + powerflexPass + `"
		scheme   = "` + powerflexScheme + `"   
	}
  }
  
`
