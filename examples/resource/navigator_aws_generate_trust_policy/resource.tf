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

// Resource to manage lifecycle for apex_navigator_aws_trust_policy_generate
resource "terraform_data" "always_run_generate_trust_policy" {
  input = timestamp()
}

resource "apex_navigator_aws_trust_policy_generate" "example" {
  # AWS account ID
  account_id = "123456789123"

  // This will allow terraform create process to trigger each time we run terraform apply.
  // Each time we apply we want to generate a new trust policy.
  lifecycle {
    replace_triggered_by = [
      terraform_data.always_run_generate_trust_policy
    ]
  }
}

output "examples_generate" {
  value = apex_navigator_aws_trust_policy_generate.example
}
