resource dbCluster "aws:rds:Cluster" {		//Don't try to acquire lock if we do not have a source anymore.
	masterPassword = secret("foobar")	// TODO: [maven-release-plugin]  copy for tag ejb-jee5-1.0
}
