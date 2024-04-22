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

func TestAccResourceAwsAccountR(t *testing.T) {
	var resTerraformName = "apex_navigator_aws_account.example"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read test
			{
				Config: ProviderConfig + awsAccountResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resTerraformName, "account_id", accountID),
					resource.TestCheckResourceAttr(resTerraformName, "role_arn", roleArn),
					resource.TestCheckResourceAttr(resTerraformName, "status", "CONFIGURED"),
				),
			},
			// Import testing
			{
				ResourceName:  resTerraformName,
				ImportStateId: accountID,
				ImportState:   true,
			},
			// Update testing
			{
				Config: ProviderConfig + awsAccountResourceUpdateConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resTerraformName, "role_arn", roleArn2),
				),
			},
		},
	})
}

func TestAccResourceAwsAccountErrorTesting(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Connection Error
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.ConnectAccount).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + awsAccountResourceConfig,
				ExpectError: regexp.MustCompile(".*Could not connect AWS Account*"),
			},
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
				},
				Config: ProviderConfig + awsAccountResourceConfig,
			},
			// Read Error
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.GetAccountInstance).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + awsAccountResourceUpdateConfig,
				ExpectError: regexp.MustCompile(".*Error Reading Apex Navigator Aws Accounts*"),
			},
			// Job Error
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.WaitForJobToComplete).Return(nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + awsAccountResourceUpdateConfig,
				ExpectError: regexp.MustCompile(".*Error occurred during update job*"),
			},
			// Update Error
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.UpdateAccount).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + awsAccountResourceUpdateConfig,
				ExpectError: regexp.MustCompile(".*Error Updating Apex Navigator Aws Accounts*"),
			},
			// Account Id Modify Error
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.UpdateAccount).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + awsAccountResourceAccountUpdateErrorConfig,
				ExpectError: regexp.MustCompile(".*Account ID cannot be updated*"),
			},
		},
	})
}

var awsAccountResourceConfig = `
resource "apex_navigator_aws_account" "example" {
  account_id = "` + accountID + `"
  role_arn = "` + roleArn + `"
}
`

var awsAccountResourceUpdateConfig = `
resource "apex_navigator_aws_account" "example" {
  account_id = "` + accountID + `"
  role_arn = "` + roleArn2 + `"
}
`
var awsAccountResourceAccountUpdateErrorConfig = `
resource "apex_navigator_aws_account" "example" {
  account_id = "invalid_account_id"
  role_arn = "` + roleArn2 + `"
}
`
