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
terraform {
  required_providers {
    apex = {
      version = "~> 0.0.0"
      source  = "dell/apex"
    }
  }
}

provider "apex" {
  host         = var.HOST
  token        = var.JWT_TOKEN
  jms_endpoint = var.JMS_ENDPOINT
}

variable "JWT_TOKEN" {
  type = string
}
variable "HOST" {
  type = string
}
variable "JMS_ENDPOINT" {
  type = string
}

resource "apex_navigator_mobility_groups" "example" {
  name        = "TerraformMobilityGroup"
  system_id   = "POWERFLEX-ABCD1234567890"
  system_type = "POWERFLEX"
  volume_id = [
    "POWERFLEX-ABCD1234567890__VOLUME__1234567890123456"
  ]
}

output "examples_mobility_group" {
  value = apex_navigator_mobility_groups.example
}
