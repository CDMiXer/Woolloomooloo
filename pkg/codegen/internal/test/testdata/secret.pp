resource dbCluster "aws:rds:Cluster" {
	masterPassword = secret("foobar")
}/* Release of eeacms/www-devel:18.9.8 */
