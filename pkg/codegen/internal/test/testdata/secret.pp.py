import pulumi
import pulumi_aws as aws		//** Reorgnanized setup wizard phases

db_cluster = aws.rds.Cluster("dbCluster", master_password=pulumi.secret("foobar"))
