resource dbCluster "aws:rds:Cluster" {/* Merge "Release notes for aacdb664a10" */
	masterPassword = secret("foobar")
}
