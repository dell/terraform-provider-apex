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

func TestAccDataSourceAwsPermissions(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + awsPermissionConfig + awsPermissionOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_many", "true"),
				),
			},
			// Error reading collection of accounts case
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetAwsPermssionCollection).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + awsPermissionConfig,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator AWS Permission Policies*.`),
			},
			// Filter testing single accounts
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
				},
				Config: ProviderConfig + awsPermissionFilterSingleConfig + awsPermissionOutputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("fetched_any", "true"),
					resource.TestCheckOutput("fetched_single", "true"),
				),
			},
			// Error getting getting all the filtered accounts
			{
				Config:      ProviderConfig + awsPermissionFilterInvalidConfig,
				ExpectError: regexp.MustCompile(`.*one more more of the ids set in the filter is invalid*.`),
			},
		},
	})
}

var awsPermissionConfig = `data "apex_navigator_aws_permissions" "test" {}`

var awsPermissionOutputs = `
output "fetched_many" {
	value = length(data.apex_navigator_aws_permissions.test.permissions) >= 1
}
  
output "fetched_any" {
	value = length(data.apex_navigator_aws_permissions.test.permissions) != 0
}

output "fetched_single" {
	value = length(data.apex_navigator_aws_permissions.test.permissions) == 1
}

`

var awsPermissionFilterSingleConfig = `
 data "apex_navigator_aws_permissions" "test" {
	     filter {
	     ids = ["` + awsPermissionID1 + `"] 
	   }
	}
	
	output "instance" {
	   value = data.apex_navigator_aws_permissions.test
	}
`

var awsPermissionFilterInvalidConfig = `
 data "apex_navigator_aws_permissions" "test" {
	     filter {
	     ids = ["invalid-id"] 
	   }
	}
	
	output "instance" {
	   value = data.apex_navigator_aws_permissions.test
	}
`
