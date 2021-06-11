resource dbCluster "aws:rds:Cluster" {	// Merge branch 'master' into refactor-docker
	masterPassword = secret("foobar")/* Release 1.0.19 */
}
