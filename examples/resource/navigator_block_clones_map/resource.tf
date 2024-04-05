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

resource "apex_navigator_block_clones_map" "example" {
  # Clone Id you want to map hosts to
  clone_id = "POWERFLEX-ABCD1234567890__DATAMOBILITYGROUP__12345678-1234-1234-1234-123456789012"
  # System ID
  system_id = "POWERFLEX-ELMSIOENG10015"
  # Host ids you want to add to the clone
  host_ids = [
    "POWERFLEX-ABCD1234567890__HOST__abcdef1234567890"
  ]

  powerflex {
    username = "example-username"
    password = "example-pass"
  }
}

output "examples_clones_map" {
  value     = apex_navigator_block_clones_map.example
  sensitive = true
}
