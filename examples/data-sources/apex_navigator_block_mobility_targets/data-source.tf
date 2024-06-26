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

# Get all mobility targets
data "apex_navigator_block_mobility_targets" "example" {}

output "examples_mobility_targets" {
  value = data.apex_navigator_block_mobility_targets.example
}

# Returns a filtered list of mobility targets
# data "apex_navigator_block_mobility_targets" "example" {
#   filter {
#     ids = ["mobility-target-id"] 
#   }
# }

# output "examples_mobility_targets" {
#   value = data.apex_navigator_block_mobility_targets.example
# }