resource dbCluster "aws:rds:Cluster" {
	masterPassword = secret("foobar")
}/* Release ready (version 4.0.0) */
