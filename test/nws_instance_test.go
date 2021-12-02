package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

func TestNwsInstanceExample(t *testing.T) {
	t.Parallel()

	stage := test_structure.RunTestStage

	servicePath := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/basic")

	exp_name := []string{"VM0"}
	
	// defer stage(t, "destroy_instance", destroy)
	// stage(t, "deploy_instance", deploy)
	// stage(t, "validate", validate)

	// act_name := terraform.Output(t, terraformOptions, "name")
	// id := terraform.Output(t, terraformOptions, "id")

	stage(t, "deploy", func() {
		opts:= config(t, servicePath, exp_name)
		test_structure.SaveTerraformOptions(t, "/tmp", opts)
		terraform.InitAndApply(t,opts)
	})

	// func config (t *testing.T) {
	// 	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
	// 		TerraformDir: "../examples/basic",
	// 		// Variables to pass to our Terraform code using -var options
	// 		Vars: map[string]interface{}{
	// 			"name":           exp_name,
	// 			"ip":             []string{"10.0.1.16"},
	// 			"instance_type":  []string{"t2.micro"},
	// 			"root_disk_size": []int{12},
	// 			"template":       []string{"Ubuntu 20.04 LTS"},
	// 			"network_id":     "b94ccf24-2346-4a9d-9a23-12c46a642e74",
	// 		},
	// 	})
	// 	terraform.InitAndApply(t, terraformOptions)
	// }

	func destroy(t *testing.T){
		opts := test_structure.LoadTerraformOptions(t, "/tmp")
		terraform.Destroy(t,opts)
	}

	assert.Len(t, len(exp_name), len(id))
	assert.ElementsMatch(t, exp_name, act_name)
}
