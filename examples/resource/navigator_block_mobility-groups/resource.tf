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

resource "apex_navigator_block_mobility_groups" "example" {
  # Name of the Mobility Group source
  name = "TerraformMobilityGroup"
  # ID of the target system
  system_id = "POWERFLEX-ABCD1234567890"
  # Type of the target system
  system_type = "POWERFLEX"
  # ID of the volume you want to add to the group
  volume_id = [
    "POWERFLEX-ABCD1234567890__VOLUME__1234567890123456"
  ]

  powerflex {
    username = "example-user"
    password = "example-pass"
    insecure = true
  }
}
output "examples_mobility_group" {
  value     = apex_navigator_block_mobility_groups.example
  sensitive = true
}
