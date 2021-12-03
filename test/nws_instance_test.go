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

	const (
		instType  = "t2.micro"
		diskSize  = 12
		template  = "Ubuntu 20.04 LTS"
		networkId = "b94ccf24-2346-4a9d-9a23-12c46a642e74"
	)

	testCases := []TestCaseT{
		{
			[]string{genVMName()},
			[]string{"10.0.1.16"},
			[]string{instType},
			[]int{diskSize},
			[]string{template},
			networkId,
		},
		{
			[]string{genVMName(), genVMName()},
			[]string{"10.0.1.17", "10.0.1.18"},
			[]string{instType, instType},
			[]int{diskSize, diskSize},
			[]string{template, template},
			networkId,
		},
	}
	for _, testCase := range testCases {

		// capture range variable so that it doesn't update when the subtest goroutine swaps.
		testCase := testCase

		stage(t, "deploy", func() {
			opts := config(t, testCase, servicePath)
			test_structure.SaveTerraformOptions(t, servicePath, opts)
			terraform.InitAndApply(t, opts)
		})

		defer stage(t, "destroy", func() {
			opts := test_structure.LoadTerraformOptions(t, servicePath)
			terraform.Destroy(t, opts)
		})

		stage(t, "validate", func() {
			opts := test_structure.LoadTerraformOptions(t, servicePath)
			validate(t, opts, testCase.name)
		})
	}
}
