package test

const (
	instType  = "t2.micro"
	diskSize  = 12
	template  = "Ubuntu 20.04 LTS"
	networkID = "202ff9b2-c751-49c9-90b9-72f7764e1a6d"
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
