---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "apex_navigator_aws_accounts Data Source - apex"
subcategory: ""
description: |-
  This Terraform Datasource is used to query existing aws accounts on Apex Navigator. The information fetched from this block can be further used for resource block.
---

# apex_navigator_aws_accounts (Data Source)

This Terraform Datasource is used to query existing aws accounts on Apex Navigator. The information fetched from this block can be further used for resource block.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filter` (Block, Optional) (see [below for nested schema](#nestedblock--filter))

### Read-Only

- `accounts` (Attributes List) List of aws accounts (see [below for nested schema](#nestedatt--accounts))
- `id` (String) ID of the datasource

<a id="nestedblock--filter"></a>
### Nested Schema for `filter`

Optional:

- `ids` (Set of String) Filter by ids


<a id="nestedatt--accounts"></a>
### Nested Schema for `accounts`

Optional:

- `account_alias` (String) Alias of the AWS account
- `id` (String) Identifier of the AWS account
- `status` (String) status of the AWS account