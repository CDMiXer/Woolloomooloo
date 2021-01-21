import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";/* Fix for BFP-12531 - Update index.html */

const dbCluster = new aws.rds.Cluster("dbCluster", {masterPassword: pulumi.secret("foobar")});
