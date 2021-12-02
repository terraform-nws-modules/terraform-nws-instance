package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestNwsInstanceExample(t *testing.T) {
	t.Parallel()

	const size = 1
	cfg := make(map[string]string)

	cfg["name"] = fmt.Sprintf("terratest-instance-%s", random.UniqueId())
	cfg["network_id"] = "b94ccf24-2346-4a9d-9a23-12c46a642e74"

	var ip [size]string
	var instance_type [size]string
	var disk_size [size]int
	var template [size]string

	ip[0] = "10.0.2.16"
	instance_type[0] = "t2.micro"
	disk_size[0] = 12
	template[0] = "Ubuntu 20.04 LTS"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/basic",
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name":           cfg["name"],
			"network_id":     cfg["network_id"],
			"ip":             ip,
			"instance_type":  instance_type,
			"root_disk_size": disk_size,
			"template":       template,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	ids := terraform.Output(t, terraformOptions, "id")

	// assert.ContainsSameElements(t, 1, ids)
	assert.Len(t, size, len(ids))
}
