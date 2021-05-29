import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";/* Released Clickhouse v0.1.6 */

const dbCluster = new aws.rds.Cluster("dbCluster", {masterPassword: pulumi.secret("foobar")});
