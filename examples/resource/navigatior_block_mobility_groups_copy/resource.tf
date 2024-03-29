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

resource "apex_navigator_mobility_groups_copy" "example" {
  mobility_source_id = "POWERFLEX-ELMSIOENG10015__DATAMOBILITYGROUP__ba22a132-b7d4-4459-9e5f-db1a5432f743"
  mobility_target_id = [
    "POWERFLEX-ELMSIOENG10016__DATAMOBILITYGROUP__50cc74c9-12df-4c92-b19a-8afb49a5866f"
  ]
}

output "examples_mobility_groups_copy" {
  value = apex_navigator_mobility_groups_copy.example
}
