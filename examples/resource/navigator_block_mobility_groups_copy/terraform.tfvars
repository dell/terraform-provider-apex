/*
Copyright (c) 2022-2024 Dell Inc., or its subsidiaries. All Rights Reserved.

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

mobility_group = {
  "group-1" = {
    powerflex_source_user     = "source_user"
    powerflex_source_password = "source_pass"
    powerflex_target_user     = "target_user"
    powerflex_target_password = "target_pass"
    insecure                  = true
    mobility_source_id        = "source-id-example-1"
    mobility_target_id        = "target-id-example-2"
  },
  "group-2" = {
    powerflex_source_user     = "source_user"
    powerflex_source_password = "source_pass"
    powerflex_target_user     = "target_user"
    powerflex_target_password = "target_pass"
    insecure                  = true
    mobility_source_id        = "source-id-example-2"
    mobility_target_id        = "target-id-example-2"
  },
}
