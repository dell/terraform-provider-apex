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
// Resource to manage lifecycle for apex_navigator_block_mobility_groups_copy
resource "terraform_data" "always_run_mobility_groups_copy" {
  input = timestamp()
}

resource "apex_navigator_block_mobility_groups_copy" "example" {
  for_each           = var.mobility_group
  mobility_source_id = each.value.mobility_source_id
  mobility_target_id = each.value.mobility_target_id

  # Note: PowerFlex credentials are required to activate the system for mobility related operations.
  powerflex_source {
    username = each.value.powerflex_source_user
    password = each.value.powerflex_source_password
    insecure = each.value.insecure
  }
  powerflex_target {
    username = each.value.powerflex_target_user
    password = each.value.powerflex_target_password
    insecure = each.value.insecure
  }
  // This will allow terraform create process to trigger each time we run terraform apply.
  lifecycle {
    replace_triggered_by = [
      terraform_data.always_run_mobility_groups_copy
    ]
  }
}

output "examples_mobility_groups_copy" {
  value     = apex_navigator_block_mobility_groups_copy.example
  sensitive = true
}
