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

data "apex_navigator_storage_products" "example" {}
// Grabs the latest file version
locals {
  latest_file_ids = [
    for k, v in data.apex_navigator_storage_products.example.storage_products :
    v.storage_type == "FILE" && v.cloud_type == "AWS" ? v.latest_version : ""
  ]
  latest_file_id = compact(local.latest_file_id)[0]
}

resource "apex_navigator_file_storage" "cloud_instance" {
  # Type of system you want to deploy
  storage_system_type = "POWERSCALE"
  # The name of the system.
  name = "apex-navigator-terraform-file"

  # This set to the latest file version. 
  # If you want to set a specific version check the apex_navigator_storage_products datasource for more versions
  product_version = local.latest_file_id
  # deployment_details (can be either system_on_prem or system_public_cloud)
  deployment_details = {
    system_public_cloud = {
      deployment_type            = "PUBLIC_CLOUD"
      cloud_type                 = "AWS"
      cloud_account              = "012345678901"
      cloud_region               = "us-east-1"
      availability_zone_topology = "SINGLE_AVAILABILITY_ZONE"
      raw_capacity               = "20"
      tier_type                  = "BALANCED"
      iam_instance_profile       = "IAMProfileTest"
      ssh_key_name               = "apex-navigator-terraform-key"
      vpc = {
        is_new_vpc = true
        vpc_name   = "my-vpc"
      }
      availability_zones = ["us-east-1a"]
      subnet_options = [
        {
          cidr_block  = "168.31.0.0/28"
          subnet_type = "EXTERNAL"
        },
        {
          cidr_block  = "10.0.0.0/22"
          subnet_type = "INTERNAL"
        }
      ]
    }
  }
  depends_on = [data.apex_navigator_storage_products.example]
}

output "navigator_file_storage" {
  sensitive = true
  value     = apex_navigator_file_storage.cloud_instance
}
