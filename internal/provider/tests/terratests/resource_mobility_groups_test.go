package tests

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// An example of how to test the simple Terraform module in examples/terraform-basic-example using Terratest.
func TestTerraformResourceMobilityGroups(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		// The path to where our Terraform code is located
		TerraformDir: "../../../../examples/tests/mobility_groups",
	})

	terraformUpdateOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		// The path to where our Terraform code is located
		TerraformDir: "../../../../examples/tests/mobility_groups",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name":        "Update",
			"description": "TestUpdate",
			"volume_id":   `["POWERFLEX-ELMVXRTEST0004__VOLUME__9e5a801700000000"]`,
		},
	})

	terraformRemoveOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		// The path to where our Terraform code is located
		TerraformDir: "../../../../examples/tests/mobility_groups/remove",
	})
	// Clean up resources with "terraform destroy". Using "defer" runs the command at the end of the test, whether the test succeeds or fails.
	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform apply".
	// This will run `terraform apply` and fail the test if there are any errors
	terraform.Apply(t, terraformOptions)

	// Run "terraform apply again"
	// This will run 'terraform apply' to update the existing state and fail the test if there are any errors
	terraform.Apply(t, terraformUpdateOptions)

	// run 'terraform apply' to remove the created mobility group resource
	terraform.Apply(t, terraformRemoveOptions)
}
