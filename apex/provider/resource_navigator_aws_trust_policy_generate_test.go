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

func TestAccResourceGenerateTrustPolicyR(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + generateAwsTrustPolicyResourceConfig,
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
			// Updates should pass since it is just a passthrough
			{
				Config:      ProviderConfig + generateAwsTrustPolicyResourceUpdateConfig,
				ExpectError: regexp.MustCompile(`.*Error generating Trust Policy*.`),
			},
		},
	})
}

func TestAccResourceGenerateTrustPolicyError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// error getting clone
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.GenerateAwsTrustPolicy).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + generateAwsTrustPolicyResourceConfig,
				ExpectError: regexp.MustCompile(`.*Error generating Trust Policy*.`),
			},
		},
	})
}

var generateAwsTrustPolicyResourceConfig = `
resource "apex_navigator_aws_trust_policy_generate" "example" {
		account_id                 = "` + accountID + `"
}
`

var generateAwsTrustPolicyResourceUpdateConfig = `
resource "apex_navigator_aws_trust_policy_generate" "example" {
		account_id                 = "some-other-id-should-fail"
}
`
