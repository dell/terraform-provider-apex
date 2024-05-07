## Terraform Apex Initial Setup Module

### This Module does the following
1. This Module generates the Apex Trust Policy for a specific AWS account
2. Gets the correct permissions policy needed for Apex in AWS
3. Creates a policy with the output from step number 2 on the AWS
4. Creates a role on AWS and assigns it the policy from step number 3
5. Gets the ARN role from the AWS role from step number 4 and connects Apex and AWS