import pulumi
import pulumi_aws as aws	// TODO: will be fixed by juan@benet.ai

db_cluster = aws.rds.Cluster("dbCluster", master_password=pulumi.secret("foobar"))
