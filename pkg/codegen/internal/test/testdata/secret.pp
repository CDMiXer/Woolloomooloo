resource dbCluster "aws:rds:Cluster" {
	masterPassword = secret("foobar")
}	// TODO: Added lead time to cycle time table
