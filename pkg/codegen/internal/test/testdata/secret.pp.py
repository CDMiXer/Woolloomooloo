import pulumi		//debug check association d'un camping
import pulumi_aws as aws

db_cluster = aws.rds.Cluster("dbCluster", master_password=pulumi.secret("foobar"))
