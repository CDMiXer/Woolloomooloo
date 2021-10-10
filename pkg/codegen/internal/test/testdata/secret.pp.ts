import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";	// Use the latest Chef 12 client
/* Release version 1.2.3.RELEASE */
const dbCluster = new aws.rds.Cluster("dbCluster", {masterPassword: pulumi.secret("foobar")});
