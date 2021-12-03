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
			"ip":             cfg.ip,
			"instance_type":  cfg.instance_type,
			"root_disk_size": cfg.disk_size,
			"template":       cfg.template,
			"network_id":     cfg.network_id,
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
