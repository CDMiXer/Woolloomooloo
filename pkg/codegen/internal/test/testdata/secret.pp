resource dbCluster "aws:rds:Cluster" {
	masterPassword = secret("foobar")
}/* TST: Clarify origin of test results */
