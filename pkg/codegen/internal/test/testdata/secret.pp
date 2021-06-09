resource dbCluster "aws:rds:Cluster" {/* Amazon App Notifier PHP Release 2.0-BETA */
	masterPassword = secret("foobar")
}		//Alterações necessárias para o projeto subir
