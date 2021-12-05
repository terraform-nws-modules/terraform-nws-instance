package test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func config(t *testing.T, cfg testCaseT, servicePath string) *terraform.Options {

	return terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: servicePath,
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name":           cfg.name,
			"ip":             cfg.ip,
			"instance_type":  cfg.instanceType,
			"root_disk_size": cfg.diskSize,
			"template":       cfg.template,
			"network_id":     networkID,
			"keypair":        keypair,
		},
		// Add retries to eliminate trasilent errors
		MaxRetries:         3,
		TimeBetweenRetries: 5 * time.Second,
	})
}

// validates Terraform output versus expected
func validate(t *testing.T, opts *terraform.Options, name []string) {
	outName := terraform.Output(t, opts, "name")
	outID := terraform.Output(t, opts, "id")

	actName := strings.Fields(trimBrackets(outName))
	actID := strings.Fields(trimBrackets(outID))

	assert.Equal(t, len(name), len(actID))
	assert.ElementsMatch(t, name, actName)
}

// Validation helpers
func trimBrackets(s string) string {
	str0 := strings.TrimLeft(s, "[")
	str1 := strings.TrimRight(str0, "]")
	return str1
}

func genVMName() string {
	return fmt.Sprintf("vm-%s", random.UniqueId())
}
