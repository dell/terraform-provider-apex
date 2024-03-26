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

Provider is in Alpha. Support is limited.

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

- [Terraform](https://www.terraform.io/downloads.html) >= 1.5.3
- [Go](https://golang.org/doc/install) >= 1.21


## List of DataSources in Terraform Provider for Dell Apex
  * [Block Storage](docs/data-sources/navigator_block_storage.md)
  * [Clone](docs/data-sources/navigator_clone.md)
  * [Hosts](docs/data-sources/navigator_hosts.md)
  * [Mobility Group](docs/data-sources/navigator_mobility_group.md)
  * [Mobility Target](docs/data-sources/navigator_mobility_target.md)
  * [Pools](docs/data-sources/navigator_pools.md)
  * [Volumes](docs/data-sources/navigator_volumes.md)

## List of Resources in Terraform Provider for Dell Apex
  * [Block Storage](docs/resources/navigator_block_storage.md)
  * [Clones Map](docs/resources/navigator_clones_map.md)
  * [Clones Refresh](docs/resources/navigator_clones_refresh.md)
  * [Clones Unmap](docs/resources/navigator_clones_unmap.md)
  * [Clones](docs/resources/navigator_clones.md)
  * [Mobility Groups Copy](docs/resources/navigator_mobility_groups_copy.md)
  * [Mobility Groups](docs/resources/navigator_mobility_groups.md)
  * [Mobility Targets](docs/resources/navigator_mobility_targets.md)

## Block Storage
Deploy and Decommision Block Storage Systems to the Cloud 

## Data Mobility
Move data on volumes between the on-premises system and the cloud
The mobility features in Dell APEX Navigator for Multicloud Storage allow you to move storage volumes between your on-premises system and the cloud.
To use the moved data, you must clone the target copy and make it accessible to hosts.
To perform mobility operations in Dell APEX Navigator, you must use a Dell Premier account that has been assigned the ITOps role.
Mobility is a feature available at no additional cost as a technology preview for a limited time.
In the future, the mobility feature may be available at an additional cost

## Installation and execution of Terraform Provider for Dell Apex
The installation and execution steps of Terraform Provider for Dell Apex can be found [here](about/INSTALLATION.md).

## Releasing, Maintenance and Deprecation

Terraform Provider for Dell Technnologies Apex follows [Semantic Versioning](https://semver.org/).

New versions will be release regularly if significant changes (bug fix or new feature) are made in the provider.

Released code versions are located on tags in the form of "vx.y.z" where x.y.z corresponds to the version number.

## Documentation

For more detailed information, please refer to [Dell Terraform Providers Documentation](https://dell.github.io/terraform-docs/).