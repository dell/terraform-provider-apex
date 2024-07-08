---
# Copyright (c) 2023-2024 Dell Inc., or its subsidiaries. All Rights Reserved.
# 
# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://mozilla.org/MPL/2.0/
# 
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

title: "apex_navigator_file_storage resource"
linkTitle: "apex_navigator_file_storage"
page_title: "apex_navigator_file_storage Resource - apex"
subcategory: ""
description: |-
  This Terraform resource is used to manage File Storage on Apex Navigator. We can create, read, update, delete File Storage on Apex Navigator.We can also import existing File Storage from Apex Navigator.
---

# apex_navigator_file_storage (Resource)

This Terraform resource is used to manage File Storage on Apex Navigator. We can create, read, update, delete File Storage on Apex Navigator.We can also import existing File Storage from Apex Navigator.

## Example Usage

### Existing VPC

```terraform
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
        is_new_vpc = false
        vpc_id     = "vpc-12345678901234567"
      }
      subnet_options = [
        {
          subnet_id   = "subnet-12345678901234567"
          subnet_type = "EXTERNAL"
        },
        {
          subnet_id   = "subnet-12345678901234567"
          subnet_type = "INTERNAL"
        },
        {
          subnet_id   = "subnet-12345678901234567"
          subnet_type = "SCG"
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
```

### New VPC

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The user-defined name of the system.
- `product_version` (String) Product version.
- `storage_system_type` (String) Type of the system

### Optional

- `bandwidth` (Number) The system bandwidth. Aggregated for a rolling average over the last 24 hours - Unit: bytes/s
- `capacity_impact` (Number) Impact point of highest impacting issue in the capacity health category
- `capacity_issue_count` (Number) Total number of issues in the capacity health category
- `compression_savings` (Number) Storage efficiency ratio of data which has compression applied to it on the system.
- `configuration_impact` (Number) Impact point of highest impacting issue in the configuration health category.
- `configuration_issue_count` (Number) Total number of issues in the configuration health category.
- `configured_size` (Number) The configured size for this system - Unit: bytes.
- `connectivity_status` (String) Connectivity status.
- `contract_coverage_type` (String) Type of the service contract of the system.
- `contract_expiration_date_timestamp` (String) Expiration date for the service contract of the system.
- `data_protection_impact` (Number) Impact point of highest impacting issue in the data protection health category.
- `data_protection_issue_count` (Number) Total number of issues in the data protection health category.
- `deployment_details` (Attributes) Details of deployment (see [below for nested schema](#nestedatt--deployment_details))
- `display_identifier` (String) Unique identifier for the system.
- `free_percent` (Number) The %free capacity.
- `free_size` (Number) The free size value - Unit: bytes.
- `health_connectivity_status` (String) Health connectivity status.
- `health_issue_count` (Number) Total amount of health issues.
- `health_score` (Number) The overall health score of the system.
- `health_state` (String) Health state of the system.
- `id` (String) Identifier for the storage system
- `iops` (Number) The IOPS for the system. Aggregated for a rolling average over the last 24 hours - Unit: IO/s.
- `ipv4_address` (String) IPv4 address of the system.
- `ipv6_address` (String) IPv6 address of the system.
- `last_contact_timestamp` (String) Last time that CloudIQ received data from the system.
- `latency` (Number) The latency for the system. Aggregated for a rolling average over the last 24 hours - Unit: microseconds.
- `license_expiration_date_timestamp` (String) Expiration date for the license on the system.
- `license_type` (String) Type of the license on the system.
- `logical_size` (Number) The logical size written - Unit: bytes.
- `model` (String) The model of the system.
- `overall_efficiency` (Number) The overall system-level storage efficiency ratio based on Thin, Snapshots, Deduplication, and Data Reduction.
- `performance_impact` (Number) Impact point of highest impacting issue in the performance health category.
- `performance_issue_count` (Number) Total number of issues in the performance health category.
- `serial_number` (String) The serial number for this system.
- `site_name` (String) Name of the site where the system is registered to.
- `snaps_savings` (Number) The snaps savings for this system.
- `system_health_impact` (Number) Health impact for the system.
- `system_health_issue_count` (Number) Total amount of health issues for the system.
- `system_id` (String) Unique identifier for the device or appliance
- `thin_savings` (Number) The savings due to thin provisioning.
- `total_size` (Number) The total size of the system - Unit: bytes.
- `unconfigured_size` (Number) The unconfigured capacity for this system - Unit: bytes.
- `used_percent` (Number) Percentage of capacity used for this system.
- `used_size` (Number) The value of used capacity for this system - Unit: bytes.
- `vendor` (String) Name of the vendor who makes the system.

### Read-Only

- `system_type` (String) Type of the system
- `version` (String) Version identifier.

<a id="nestedatt--deployment_details"></a>
### Nested Schema for `deployment_details`

Optional:

- `system_on_prem` (Attributes) (see [below for nested schema](#nestedatt--deployment_details--system_on_prem))
- `system_public_cloud` (Attributes) (see [below for nested schema](#nestedatt--deployment_details--system_public_cloud))

<a id="nestedatt--deployment_details--system_on_prem"></a>
### Nested Schema for `deployment_details.system_on_prem`

Optional:

- `city` (String) Name of the city where the system is located.
- `country` (String) Name of the country where the system is located.
- `deployment_type` (String) System deployment types (e.g. PUBLIC_CLOUD)
- `location` (String) User provided description of where the system is located.
- `site_name` (String) Name of the site where the system is located
- `state` (String) Name of the state where the system is located.
- `street_address_1` (String) Street address 1 of where the system is located
- `street_address_2` (String) Street address 2 of where the system is located
- `zip_code` (String) Zip code of where the system is located


<a id="nestedatt--deployment_details--system_public_cloud"></a>
### Nested Schema for `deployment_details.system_public_cloud`

Required:

- `cloud_account` (String) Cloud provider account where the storage system resides
- `cloud_region` (String) Cloud provider region where the storage system resides
- `cloud_type` (String) Enum for all the supported cloud providers * AWS - Amazon Web Services
- `iam_instance_profile` (String) IAM instance profile requested during the deployment time.This Role should already be created on the IAM instance.
- `raw_capacity` (String) Raw capacity requested during the deployment time - Unit: bytes
- `ssh_key_name` (String)
- `subnet_options` (Attributes List) Subnet options (see [below for nested schema](#nestedatt--deployment_details--system_public_cloud--subnet_options))
- `tier_type` (String) Tier type requested during the deployment time
- `vpc` (Attributes) VPC (see [below for nested schema](#nestedatt--deployment_details--system_public_cloud--vpc))

Optional:

- `availability_zone_topology` (String) This enum represents all the availability zone topology * SINGLE_AVAILABILITY_ZONE * MULTIPLE_AVAILABILITY_ZONE
- `availability_zones` (List of String) For deployments in an existing VPC, this option is not required as APEX Navigator deploys the storage system to the availability zone of the subnets provided in the subnet_options. For deployments in a new VPC, APEX Navigator deploys APEX File Storage to this specific availability zone.
- `cloud_management_address` (String) Management IPv4 or IPv6 address or DNS name of the storage system in cloud
- `deployment_type` (String) System deployment types (e.g. PUBLIC_CLOUD)
- `minimum_capacity` (Number) Minimum capacity requested during the deployment time - Unit: bytes
- `minimum_iops` (Number) Minimum IOPS requested during the deployment time - Unit: IO/s
- `virtual_private_cloud` (String) Cloud virtual private environment identifier

<a id="nestedatt--deployment_details--system_public_cloud--subnet_options"></a>
### Nested Schema for `deployment_details.system_public_cloud.subnet_options`

Optional:

- `cidr_block` (String) CIDR block for subnet. It is not supported when existing vpc is used to create block storage
- `subnet_id` (String)
- `subnet_type` (String) subnet type is one of EXTERNAL, INTERNAL, SCG


<a id="nestedatt--deployment_details--system_public_cloud--vpc"></a>
### Nested Schema for `deployment_details.system_public_cloud.vpc`

Optional:

- `is_new_vpc` (Boolean)
- `vpc_id` (String)
- `vpc_name` (String)

## Import

Import is supported using the following syntax:

```shell
# /*
# Copyright (c) 2021-2024 Dell Inc., or its subsidiaries. All Rights Reserved.

# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://mozilla.org/MPL/2.0/


# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# */

terraform import apex_navigator_file_storage.cloud_instance "{\"system_id\":\"ONEFS-0affd63a92239cfd3466f920a9eb6f578293\",\"username\":\"admin\",\"password\":\"Admin123!\",\"host\":\"10.0.1.14:8080\",\"insecure\":true,\"scheme\":\"https\"}"
```