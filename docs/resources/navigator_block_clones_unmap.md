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

title: "apex_navigator_block_clones_unmap resource"
linkTitle: "apex_navigator_block_clones_unmap"
page_title: "apex_navigator_block_clones_unmap Resource - apex"
subcategory: ""
description: |-
  This Terraform resource is used to unmap hosts from clones.
---

# apex_navigator_block_clones_unmap (Resource)

This Terraform resource is used to unmap hosts from clones.


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
// Resource to manage lifecycle for apex_navigator_block_clones_unmap
resource "terraform_data" "always_run_mobility_clones_unmap" {
  input = timestamp()
}

resource "apex_navigator_block_clones_unmap" "example" {
  # Clone Id you want to unmap hosts from
  clone_id = "POWERFLEX-ABCD1234567890__DATAMOBILITYGROUP__12345678-1234-1234-1234-123456789012"
  # Host ids(these are ids generated by APEX Navigator after you map host) you want to remove from the clone
  host_ids = [
    "POWERFLEX-ABCD1234567890__DATAMOBILITYHOST__12345678-1234-1234-1234-123456789012"
  ]
  # Note: PowerFlex credentials are required to activate the system for clones related operations.
  powerflex {
    username = "example-username"
    password = "example-pass"
  }

  // This will allow terraform create process to trigger each time we run terraform apply.
  lifecycle {
    replace_triggered_by = [
      terraform_data.always_run_mobility_clones_unmap
    ]
  }
}

output "examples_clones_unmap" {
  value     = apex_navigator_block_clones_unmap.example
  sensitive = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `clone_id` (String) The ID of the clone
- `host_ids` (List of String) List of ids(generated by APEX Navigator) for host to unmap from the clone
- `system_id` (String)

### Optional

- `host_mappings` (Attributes List) This represents the mapping of a repurposed (clone) storage object to a host (presumably using the clone for some analytical workload) (see [below for nested schema](#nestedatt--host_mappings))
- `powerflex` (Block, Optional) (see [below for nested schema](#nestedblock--powerflex))
- `status` (String) Status of the clone map Job

### Read-Only

- `id` (String) ID of the clone map job

<a id="nestedatt--host_mappings"></a>
### Nested Schema for `host_mappings`

Read-Only:

- `host_id` (String) Identifier of the host
- `host_initiator_protocol` (String) Type of the host
- `host_ip` (String) IP address of host
- `host_name` (String) Name of host/SDC to be mapped to the clone
- `id` (String) This is a host mappings id generated by APEX Navigator for Multicloud Storage
- `nqn` (String) NVMe qualified name. Only applicable if host_initiator_protocol is NVMe


<a id="nestedblock--powerflex"></a>
### Nested Schema for `powerflex`

Required:

- `password` (String, Sensitive) Password of the powerflex
- `username` (String) Username of the powerflex

Optional:

- `host` (String) Host, ip or hostname of the powerflex. If left empty we will attempt to get the ip through Apex from the ID
- `insecure` (Boolean) Validated the certificate when connecting to the powerflex, defaults if unset to true
- `scheme` (String) Scheme of the powerflex, defaults if unset to https

