package test

import (
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

type TestCaseT struct {
	name          []string
	ip            []string
	instance_type []string
	disk_size     []int
	template      []string
	network_id    string
}

func config(t *testing.T, cfg TestCaseT, servicePath string) *terraform.Options {

	return terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: servicePath,
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name":           cfg.name,
			"ip":             []string{"10.0.1.16", "10.0.1.17"},
			"instance_type":  []string{"t2.micro", "t2.micro"},
			"root_disk_size": []int{12, 12},
			"template":       []string{"Ubuntu 20.04 LTS", "Ubuntu 20.04 LTS"},
			"network_id":     "b94ccf24-2346-4a9d-9a23-12c46a642e74",
		},
		MaxRetries:         3,
		TimeBetweenRetries: 5 * time.Second,
	})
}

func validate(t *testing.T, opts *terraform.Options, name []string) {
	out_name := terraform.Output(t, opts, "name")
	out_id := terraform.Output(t, opts, "id")

	act_name := strings.Fields(trimBrackets(out_name))
	act_id := strings.Fields(trimBrackets(out_id))

	assert.Equal(t, len(name), len(act_id))
	assert.ElementsMatch(t, name, act_name)
}

func trimBrackets(s string) string {
	str0 := strings.TrimLeft(s, "[")
	str1 := strings.TrimRight(str0, "]")
	return str1
}
