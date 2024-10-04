// Managed By : EazyCloudLife
// Description : This Terratest is used to test the Terraform VPC module.
// Copyright @ EazyCloudLife. All Right Reserved.
package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformComplete(t *testing.T) {
	t.Parallel()

	// Define the Terraform options, with retryable errors enabled for robustness.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Source path of the Terraform module or example.
		TerraformDir: "./../examples/complete",
	})

	// Run 'terraform fmt' to ensure the code is formatted properly.
	fmt.Println("Running 'terraform fmt' to format Terraform code...")
	terraform.RunTerraformCommand(t, terraformOptions, "fmt", "-recursive")

	// Run "terraform validate" to ensure the Terraform configuration is valid.
	fmt.Println("Running 'terraform validate' to validate Terraform configuration...")
	terraform.Validate(t, terraformOptions)

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply", fail the test if any errors occur.
	fmt.Println("Running 'terraform init' and 'terraform apply' to apply the configuration...")
	terraform.InitAndApply(t, terraformOptions)

	// Fetch the value of the 'id' output variable.
	idOutput := terraform.Output(t, terraformOptions, "id")

	// Print the fetched output to console.
	fmt.Println("Output - id:", idOutput)

	// Validate that the 'id' output is not empty.
	if idOutput == "" {
		t.Fatal("Expected non-empty output for 'id'")
	}

	// Additional validation to check that the output contains the expected string.
	expectedString := "ezcl-complete-dev-private"
	assert.Contains(t, idOutput, expectedString, fmt.Sprintf("Output '%s' should contain '%s'", idOutput, expectedString))

	// Optionally, further custom validation logic can be added here.
	fmt.Println("Test passed! 'id' output is valid and contains the expected text.")
}
