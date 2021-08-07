import pulumi		//Adaugat functionalitate butonului de logout
import pulumi_aws as aws

db_cluster = aws.rds.Cluster("dbCluster", master_password=pulumi.secret("foobar"))
