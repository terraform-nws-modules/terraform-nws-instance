package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func config(t *testing.T, name []string, servicePath string) *terraform.Options {
	return terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: servicePath,
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name":           name,
			"ip":             []string{"10.0.1.16"},
			"instance_type":  []string{"t2.micro"},
			"root_disk_size": []int{12},
			"template":       []string{"Ubuntu 20.04 LTS"},
			"network_id":     "b94ccf24-2346-4a9d-9a23-12c46a642e74",
		},
	})
}

func validate(t *testing.T, opts *terraform.Options, name []string) {
	act_name := terraform.Output(t, opts, "name")
	id := terraform.Output(t, opts, "id")

	assert.Len(t, len(name), len(id))
	assert.ElementsMatch(t, name, act_name)
}