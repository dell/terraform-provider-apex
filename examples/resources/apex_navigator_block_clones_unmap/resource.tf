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
// Resource to manage lifecycle for apex_navigator_block_clones_unmap
resource "terraform_data" "always_run_mobility_clones_unmap" {
  input = timestamp()
}

resource "apex_navigator_block_clones_unmap" "example" {
  # Clone Id you want to unmap hosts from
  clone_id = "POWERFLEX-ABCD1234567890__DATAMOBILITYGROUP__12345678-1234-1234-1234-123456789012"
  # Host ids(these are ids generated by APEX Navigator after you map host) you want to remove from the clone
  host_ids = [
    "POWERFLEX-ABCD1234567890__DATAMOBILITYHOST__12345678-1234-1234-1234-123456789012"
  ]
  # Note: PowerFlex credentials are required to activate the system for clones related operations.
  powerflex {
    username = "example-username"
    password = "example-pass"
  }

  // This will allow terraform create process to trigger each time we run terraform apply.
  lifecycle {
    replace_triggered_by = [
      terraform_data.always_run_mobility_clones_unmap
    ]
  }
}

output "examples_clones_unmap" {
  value     = apex_navigator_block_clones_unmap.example
  sensitive = true
}