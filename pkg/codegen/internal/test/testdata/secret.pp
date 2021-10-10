resource dbCluster "aws:rds:Cluster" {
	masterPassword = secret("foobar")/* Merge "Gerrit 2.4 ReleaseNotes" into stable-2.4 */
}
