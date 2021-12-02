package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestNwsInstanceExample(t *testing.T) {
	t.Parallel()

	stage := test_structure.RunTestStage

	servicePath := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/basic")

	exp_name := []string{"VM0", "VM1"}

	stage(t, "deploy", func() {
		opts := config(t, exp_name, servicePath)
		test_structure.SaveTerraformOptions(t, servicePath, opts)
		terraform.InitAndApply(t, opts)
	})

	defer stage(t, "destroy", func() {
		opts := test_structure.LoadTerraformOptions(t, servicePath)
		terraform.Destroy(t, opts)
	})

	stage(t, "validate", func() {
		opts := test_structure.LoadTerraformOptions(t, servicePath)
		validate(t, opts, exp_name)
	})

}
