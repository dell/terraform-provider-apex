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

resource "apex_navigator_block_mobility_targets" "example" {
  # Name of the Mobility target
  name = "TerraformMobilityTarget"
  # Source Mobility Group Id
  source_mobility_group_id = "POWERFLEX-ABCD1234567890__DATAMOBILITYGROUP__12345678-1234-1234-1234-123456789012"
  # Target System Id
  system_id = "POWERFLEX-ABCD1234567890"
  # Target System Type
  system_type = "POWERFLEX"
  # Storage pool id to use for allocating target volumes
  target_system_options = "POWERFLEX-ABCD1234567890_STORAGE_POOL__1234567890123456"

  # The source mobility group Powerflex
  powerflex_source {
    username = "example-source-username"
    password = "example-source-pass"
  }

  # The Powerflex where you want to create the target
  powerflex_target {
    username = "example-target-username"
    password = "example-target-pass"
  }
}

output "examples_mobility_target" {
  value     = apex_navigator_block_mobility_targets.example
  sensitive = true
}
