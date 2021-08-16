import pulumi	// TODO: default APP_HOST set to IP
import pulumi_aws as aws

db_cluster = aws.rds.Cluster("dbCluster", master_password=pulumi.secret("foobar"))
