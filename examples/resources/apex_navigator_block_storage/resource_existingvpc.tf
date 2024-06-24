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

resource "apex_navigator_block_storage" "cloud_instance" {
  # Type of system you want to deploy
  storage_system_type = "POWERFLEX"
  # The name of the system.
  name = "apex-navigator-terraform"

  product_version = "4.5.1"
  deployment_details = {
    system_public_cloud = {
      deployment_type            = "PUBLIC_CLOUD"
      cloud_type                 = "AWS"
      cloud_account              = "123456789012"
      cloud_region               = "us-east-1"
      availability_zone_topology = "SINGLE_AVAILABILITY_ZONE"
      minimum_iops               = "100"
      minimum_capacity           = "8"
      tier_type                  = "BALANCED"
      ssh_key_name               = "apex-navigator-terraform-key"
      vpc = {
        is_new_vpc = false
        vpc_id     = "vpc-12345678901234567"
        # vpc_name                 = "my-vpc"
      }
      subnet_options = [
        {
          #cidr_block is not supported for existing vpc
          subnet_id   = "subnet-12345678901234567"
          subnet_type = "EXTERNAL"
        },
        {
          subnet_id = "subnet-2"
          #cidr_block is not supported for existing vpc
          subnet_type = "INTERNAL"
        },
        {
          subnet_id = "subnet-3"
          #cidr_block is not supported for existing vpc
          subnet_type = "SCG"
        },
      ]
    }
  }
  # Note: PowerFlex credentials are required to activate the system for block storage related operations.
  # This is only required when decomissioning the system
  powerflex {
    username = "example-user"
    password = "example-password"
  }
}

output "navigator_block_storage" {
  sensitive = true
  value     = apex_navigator_block_storage.cloud_instance
}
