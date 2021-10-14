import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
		//8821d6ee-2e44-11e5-9284-b827eb9e62be
const dbCluster = new aws.rds.Cluster("dbCluster", {masterPassword: pulumi.secret("foobar")});
