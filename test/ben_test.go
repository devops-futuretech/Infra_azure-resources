package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Test case to apply and validate Terraform configuration
func TestTerraformWithTerratest(t *testing.T) {
	t.Parallel()

	// 1. Set random string for unique resource names
	// uniqueID := random.UniqueId()
	terraformOptions := &terraform.Options{
		// Path to the Terraform code we will test
		TerraformDir: "../environments/dev",

		// // Variables to pass to our Terraform configuration
		// Vars: map[string]interface{}{
		// 	"unique_id": uniqueID,
		// },

		// Variables to pass to our Terraform configuration (optional)
		VarFiles: []string{
			"../../environments/dev/terraform.tfvars",
		},

		// Disable colors in Terraform commands so its easier to read stdout/stderr
		NoColor: true,
	}

	// 2. Run terraform init (initialize Terraform configuration)
	terraform.Init(t, terraformOptions)

	// 3. Run terraform apply (apply the Terraform configuration)
	terraform.Apply(t, terraformOptions)

	// 4. Get outputs from Terraform
	outputValue := terraform.Output(t, terraformOptions, "output_name")

	// 5. Validate the output value
	assert.Equal(t, "expected_output", outputValue)

	// 6. Check if the resource is created by checking resource ID or any specific attribute
	resourceID := terraform.Output(t, terraformOptions, "resource_id")
	assert.NotEmpty(t, resourceID)

	// Print the output to see in the logs
	fmt.Println("Resource ID: ", resourceID)
	fmt.Println("Output Value: ", outputValue)

	// 7. Run terraform destroy (clean up after tests)
	terraform.Destroy(t, terraformOptions)
}
