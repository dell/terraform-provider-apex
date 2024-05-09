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



variable "HOST" {
  type    = string
  default = "https://example-host.com"
}

variable "JMS_ENDPOINT" {
  type    = string
  default = "https://example-jobs-endpoint.com"
}

variable "mobility_group" {
  type = map(object({
    powerflex_source_user     = string
    powerflex_source_password = string
    powerflex_target_user     = string
    powerflex_target_password = string
    insecure                  = bool
    mobility_source_id        = string
    mobility_target_id        = string
  }))
}
