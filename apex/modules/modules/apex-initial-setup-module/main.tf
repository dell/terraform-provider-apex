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

data "apex_navigator_aws_permissions" "aws_permission" {}

resource "apex_navigator_aws_trust_policy_generate" "generated_policy" {
  # AWS account ID
  account_id = var.aws_account_id
}

resource "aws_iam_role" "test_role" {
  name = var.role_name

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  assume_role_policy = jsonencode({
    Version = aws_permission.aws_permission.version
    # for_each = {
    #   aws_permission.aws_permission.statements
    # }
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      },
    ]
  })
   depends_on = [ apex_navigator_aws_trust_policy_generate.generated_policy, data.apex_navigator_aws_permissions.aws_permission ]
}
