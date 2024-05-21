<!--
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
-->
# Terraform Provider Apex Navigator

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0%20adopted-ff69b4.svg)](about/CODE_OF_CONDUCT.md)
[![License](https://img.shields.io/badge/License-MPL_2.0-blue.svg)](LICENSE)

The Terraform Provider for Dell Technologies (Dell) Apex allows Data Center and IT administrators to use Hashicorp Terraform to automate and orchestrate the provisioning and management of Dell Apex resources

Provider is in Beta. Support is limited.

## Table of contents

* [Code of Conduct](https://github.com/dell/dell-terraform-providers/blob/main/docs/CODE_OF_CONDUCT.md)
* [Maintainer Guide](https://github.com/dell/dell-terraform-providers/blob/main/docs/MAINTAINER_GUIDE.md)
* [Committer Guide](https://github.com/dell/dell-terraform-providers/blob/main/docs/COMMITTER_GUIDE.md)
* [Contributing Guide](https://github.com/dell/dell-terraform-providers/blob/main/docs/CONTRIBUTING.md)
* [Support](#support)
* [License](#license)
* [Prerequisites](#Prerequisites)
* [List of DataSources in Terraform Provider for Dell Apex](#list-of-datasources-in-terraform-provider-for-dell-apex)
* [List of Resources in Terraform Provider for Dell Apx](#list-of-resources-in-terraform-provider-for-dell-apex)
* [Releasing, Maintenance and Deprecation](#releasing-maintenance-and-deprecation)
* [Documentation](#documentation)

## Support
For any Terraform Provider for Dell Apex issues, questions or feedback, please follow our [support process](https://github.com/dell/dell-terraform-providers/blob/main/docs/SUPPORT.md)

## License
The Terraform Provider for Dell Apex is released and licensed under the MPL-2.0 license. See [LICENSE](LICENSE) for the full terms.

## Prerequisites

- [Apex_Auth_Scripts](https://github.com/dell/terraform-provider-apex/blob/main/scripts/saml_script/README.md) This must be configured ahead of running any terraform commands in order to authenticate with Apex.

| **Terraform Provider**  | **OS**                    | **Terraform**               | **Golang** |
|------------------------ |:-----------------------   |:--------------------------  |------------|
| v1.0.0-beta             | ubuntu22.04 <br>  rhel9.x | 1.5.x <br> 1.6.x            | 1.21       |


## List of DataSources in Terraform Provider for Dell Apex
  * [AWS Accounts](docs\data-sources\navigator_aws_accounts.md)
  * [AWS Permissions](docs\data-sources\navigator_aws_permissions.md)
  * [Storages](docs\data-sources\navigator_storages.md)
  * [Block Clones](docs\data-sources\navigator_block_clones.md)
  * [Block Hosts](docs\data-sources\navigator_block_hosts.md)
  * [Block Mobility Groups](docs\data-sources\navigator_block_mobility_groups.md)
  * [Block Mobility Targets](docs\data-sources\navigator_block_mobility_targets.md)
  * [Block Pools](docs\data-sources\navigator_block_pools.md)
  * [Storage Products](docs\data-sources\navigator_storage_products.md)
  * [Block Volumes](docs\data-sources\navigator_block_volumes.md)

## List of Resources in Terraform Provider for Dell Apex
  * [AWS Account](docs\resources\navigator_aws_account.md)
  * [AWS Trust Policy Generate](docs\resources\navigator_aws_trust_policy_generate.md)
  * [Block Storage](docs\resources\navigator_block_storage.md)
  * [Block Clones](docs\resources\navigator_block_clones.md)
  * [Block Clones Map](docs\resources\navigator_block_clones_map.md)
  * [Block Clones Refresh](docs\resources\navigator_block_clones_refresh.md)
  * [Block Clones Unmap](docs\resources\navigator_block_clones_unmap.md)
  * [Block Mobility Groups Copy](docs\resources\navigator_block_mobility_groups_copy.md)
  * [Block Mobility Groups](docs\resources\navigator_block_mobility_groups.md)
  * [Block Mobility Targets](docs\resources\navigator_block_mobility_targets.md)
  * [File Storage](docs/resources/navigator_file_storage.md)

## Installation and execution of Terraform Provider for Dell Apex
The installation and execution steps of Terraform Provider for Dell Apex can be found [here](about/INSTALLATION.md).

## Other useful provider information
Once you have utilized the Terraform APEX Provider for the deployment of APEX Block Storage and/or APEX File Storage in the public cloud, subsequent automation of block and file storage workflows via the respective storage end-points can be achieved by employing the [Terraform PowerFlex Provider](https://registry.terraform.io/providers/dell/powerflex/latest) and/or the [Terraform PowerScale Provider](https://registry.terraform.io/providers/dell/powerscale/latest).

## Releasing, Maintenance and Deprecation

Terraform Provider for Dell Technnologies Apex follows [Semantic Versioning](https://semver.org/).

New versions will be release regularly if significant changes (bug fix or new feature) are made in the provider.

Released code versions are located on tags in the form of "vx.y.z" where x.y.z corresponds to the version number.

## Documentation

For more detailed information, please refer to [Dell Terraform Providers Documentation](https://dell.github.io/terraform-docs/).