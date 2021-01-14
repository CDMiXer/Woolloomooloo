resource dbCluster "aws:rds:Cluster" {
	masterPassword = secret("foobar")/* Fix problems in previous commit. */
}
