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

/*
The idea of this script is to make it easier to gather the real test ids for acceptance testing in the apex provider
It will use the get full collections for each datasource, the provide 2 IDs which can be used in the apex.env file to run tests against the actual apex instance
*/

terraform {
  required_providers {
    apex = {
      source = "dell/apex"
    }
  }
}

provider "apex" {
  host  = var.HOST
  token = var.JWT_TOKEN
}

variable "JWT_TOKEN" {
  type    = string
  default = "insert-token-here"
}

variable "HOST" {
  type    = string
  default = "https://example-host"
}

// Volumes
data "apex_navigator_block_volumes" "example" {}

output "VOLUME_ID1" {
  value = data.apex_navigator_block_volumes.example.volumes[0].id
}

output "VOLUME_ID2" {
  value = data.apex_navigator_block_volumes.example.volumes[1].id
}

// Clones
data "apex_navigator_block_clones" "example" {}

output "CLONE_ID1" {
  value = data.apex_navigator_block_clones.example.clones[0].id
}

output "CLONE_ID2" {
  value = data.apex_navigator_block_clones.example.clones[1].id
}

// Block Storage
data "apex_navigator_storages" "example" {}

output "BLOCK_STORAGE_ID_1" {
  value = data.apex_navigator_storages.example.block_storages[0].id 
}

output "BLOCK_STORAGE_ID_2" {
  value = data.apex_navigator_storages.example.block_storages[1].id
}

// Mobility Groups
data "apex_navigator_block_mobility_groups" "example" {}

output "MOBILITY_ID_1" {
  value = data.apex_navigator_block_mobility_groups.example.mobility_groups[0].id
}

output "MOBILITY_ID_2" {
  value = data.apex_navigator_block_mobility_groups.example.mobility_groups[1].id
}

// Mobility targets
data "apex_navigator_block_mobility_targets" "example" {}

output "MOBILITY_TARGET_ID_1" {
  value = data.apex_navigator_block_mobility_targets.example.mobility_targets[0].id
}

output "MOBILITY_TARGET_ID_2" {
  value = data.apex_navigator_block_mobility_targets.example.mobility_targets[1].id
}

// Hosts
data "apex_navigator_block_hosts" "example" {}

output "HOST_ID_1" {
  value = data.apex_navigator_block_hosts.example.hosts[0].id
}

output "HOST_ID_2" {
  value = data.apex_navigator_block_hosts.example.hosts[1].id
}

// Pools
data "apex_navigator_block_pools" "example" {}

output "SOURCE_POOL_ID1" {
  value = data.apex_navigator_block_pools.example.pools[0].id
}

output "SOURCE_POOL_ID2" {
  value = data.apex_navigator_block_pools.example.pools[1].id
}

// AWS Account
data "apex_navigator_aws_accounts" "example" {}

output "AWS_ACCOUNT_ID1" {
  value = data.apex_navigator_aws_accounts.example.account[0].id
}

output "AWS_ACCOUNT_ID2" {
  value = data.apex_navigator_aws_accounts.example.account[1].id
}

// File Storage
data "apex_navigator_file_storages" "example" {}

output "FILE_STORAGE_ID_1" {
  value = data.apex_navigator_file_storages.example.file_storages[0].id 
}

output "FILE_STORAGE_ID_2" {
  value = data.apex_navigator_file_storages.example.file_storages[1].id
}