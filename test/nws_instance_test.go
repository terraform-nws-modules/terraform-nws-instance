package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestNwsInstanceExample(t *testing.T) {
	t.Parallel()

	stage := test_structure.RunTestStage
	servicePath := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/basic")

	testCases := []TestCaseT{
		{
			[]string{fmt.Sprintf("vm-%s", random.UniqueId())},
			[]string{"10.0.1.16"},
			[]string{"t2.micro"},
			[]int{12},
			[]string{"Ubuntu 20.04 LTS"},
			"b94ccf24-2346-4a9d-9a23-12c46a642e74",
		},
		{
			[]string{fmt.Sprintf("vm-%s", random.UniqueId()), fmt.Sprintf("vm-%s", random.UniqueId())},
			[]string{"10.0.1.17", "10.0.1.18"},
			[]string{"t2.micro", "t2.micro"},
			[]int{12, 12},
			[]string{"Ubuntu 20.04 LTS", "Ubuntu 20.04 LTS"},
			"b94ccf24-2346-4a9d-9a23-12c46a642e74",
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
