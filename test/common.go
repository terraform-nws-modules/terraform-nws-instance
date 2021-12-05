package test

const (
	instType  = "t2.micro"
	diskSize  = 12
	template  = "Ubuntu 20.04 LTS"
	networkID = "aa1bd9e4-c308-46b9-b52e-28b16f6efc32"
	keypair   = "bku"
)

type testCaseT struct {
	testName     string
	name         []string
	ip           []string
	instanceType []string
	diskSize     []int
	template     []string
}
