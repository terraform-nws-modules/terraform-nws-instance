package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestNwsInstanceExample(t *testing.T) {
	t.Parallel()

	cfg := make(map[string]string)
	name := []string{"VM0"}

	cfg["network_id"] = "b94ccf24-2346-4a9d-9a23-12c46a642e74"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/basic",
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name":           name,
			"ip":             []string{"10.0.1.16"},
			"instance_type":  []string{"t2.micro"},
			"root_disk_size": []int{12},
			"template":       []string{"Ubuntu 20.04 LTS"},
			"network_id":     cfg["network_id"],
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	act_name := terraform.Output(t, terraformOptions, "name")
	id := terraform.Output(t, terraformOptions, "id")

	assert.Len(t, len(name), len(id))
	assert.ElementsMatch(t, name, act_name)
}
