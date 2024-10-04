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

func TestTerraformHelloWorld(t *testing.T) {
	t.Parallel()

	// Define the Terraform options, with retryable errors enabled for robustness.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Source path of the Terraform module or example.
		TerraformDir: "./../examples/hello-world",
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

	// Fetch the value of the 'hello_world' output variable.
	helloWorldOutput := terraform.Output(t, terraformOptions, "hello_world")

	// Print the fetched output to console.
	fmt.Println("Output - hello_world:", helloWorldOutput)

	// Validate that the 'hello_world' output is not empty.
	if helloWorldOutput == "" {
		t.Fatal("Expected non-empty output for 'hello_world'")
	}

	// Additional validation to check that the output contains the expected string.
	expectedString := "Hello, World"
	assert.Contains(t, helloWorldOutput, expectedString, fmt.Sprintf("Output '%s' should contain '%s'", helloWorldOutput, expectedString))

	// Optionally, further custom validation logic can be added here.
	fmt.Println("Test passed! 'hello_world' output is valid and contains the expected text.")
}
