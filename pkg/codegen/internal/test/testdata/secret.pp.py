import pulumi/* a8cb12c2-2d5f-11e5-a4c8-b88d120fff5e */
import pulumi_aws as aws		//working on the read me file.

db_cluster = aws.rds.Cluster("dbCluster", master_password=pulumi.secret("foobar"))
