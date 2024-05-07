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

// Needed for APEX AWS initial setup module
provider "aws" {
  region = "us-east-2"
}

provider "apex" {
  // required
  host         = var.HOST
  // Optional if the JMS Endpoint is different then the Host
  jms_endpoint = var.JMS_ENDPOINT
  // Optional if you are using a manual JWT token
  token        = var.JWT_TOKEN
}

module "apex-aws-initial-setup" {
  source = "./modules/apex-aws-initial-setup"
  aws_account = var.aws_account_id
  policy_name = var.policy_name
  role_name = var.role_name
}