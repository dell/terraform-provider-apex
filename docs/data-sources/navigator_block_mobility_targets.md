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

title: "apex_navigator_block_mobility_targets data source"
linkTitle: "apex_navigator_block_mobility_targets"
page_title: "apex_navigator_block_mobility_targets Data Source - apex"
subcategory: ""
description: |-
  This Terraform Datasource is used to query existing mobility targets on Apex Navigator. The information fetched from this block can be further used for resource block.
---

# apex_navigator_block_mobility_targets (Data Source)

This Terraform Datasource is used to query existing mobility targets on Apex Navigator. The information fetched from this block can be further used for resource block.

## Example Usage

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

# Get all mobility targets
data "apex_navigator_block_mobility_targets" "example" {}

output "examples_mobility_targets" {
  value = data.apex_navigator_block_mobility_targets.example
}

# Returns a filtered list of mobility targets
# data "apex_navigator_block_mobility_targets" "example" {
#   filter {
#     ids = ["mobility-target-id"] 
#   }
# }

# output "examples_mobility_targets" {
#   value = data.apex_navigator_block_mobility_targets.example
# }
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filter` (Block, Optional) (see [below for nested schema](#nestedblock--filter))

### Read-Only

- `id` (String) ID of the Mobility Target datasource
- `mobility_targets` (Attributes List) List of Mobility Targets (see [below for nested schema](#nestedatt--mobility_targets))

<a id="nestedblock--filter"></a>
### Nested Schema for `filter`

Optional:

- `ids` (Set of String) Filter by ids


<a id="nestedatt--mobility_targets"></a>
### Nested Schema for `mobility_targets`

Optional:

- `id` (String) Idenifier of this target mobility group
- `type` (String)

Read-Only:

- `bandwidth_limit` (Number) Bandwidth limit in Mbps (Mega bits per second).
- `creation_timestamp` (String) Timestamp from when the group was created
- `description` (String) Description of the mobility target
- `image_timestamp` (String) Timestamp of the last source image copied to this target
- `last_copy_job_id` (String) Last copy job ID
- `name` (String) Name of the mobility target
- `source_mobility_group_id` (String) ID of the source mobility group
- `system_id` (String) ID of the target system
- `system_type` (String) The source system type (e.g.: POWERFLEX)
- `target_members` (Attributes List) A mobility member map is a mapping of a mobility member and it's related member. Ex: a target volume with a reference to the source volume or a clone volume and its related target volume. (see [below for nested schema](#nestedatt--mobility_targets--target_members))
- `target_system_options` (String) Storage pool id to use for allocating target volumes

<a id="nestedatt--mobility_targets--target_members"></a>
### Nested Schema for `mobility_targets.target_members`

Read-Only:

- `id` (String) ID of the member
- `name` (String) Name of the member
- `parent_id` (String) Identifier of the related mobility member
- `size` (String) Size of the member