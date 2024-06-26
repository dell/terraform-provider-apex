---
# Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.
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

listIgnoreTitle: true
weight: 1
title: "apex provider"
linkTitle: "Provider"
page_title: "apex Provider"
subcategory: ""
description: |-
  The Terraform provider for Dell APEX Navigator.
---

# apex Provider

The Terraform provider for Dell APEX Navigator.

## Example Usage

provider.tf
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

provider "apex" {
  host         = var.HOST
  jms_endpoint = var.JMS_ENDPOINT
  insecure     = true
}



variable "HOST" {
  type    = string
  default = "https://apex.apis.dell.com/apex"
}

variable "JMS_ENDPOINT" {
  type    = string
  default = "https://apex.apis.dell.com/apex"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `host` (String)
- `insecure` (Boolean) Boolean variable to specify whether to validate SSL certificate or not.
- `jms_endpoint` (String)
- `token` (String, Sensitive)