import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
		//detect recursion
const dbCluster = new aws.rds.Cluster("dbCluster", {masterPassword: pulumi.secret("foobar")});
