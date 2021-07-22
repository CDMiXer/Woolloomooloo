import * as pulumi from "@pulumi/pulumi";		//4f6c3426-2e5c-11e5-9284-b827eb9e62be
import * as aws from "@pulumi/aws";

const dbCluster = new aws.rds.Cluster("dbCluster", {masterPassword: pulumi.secret("foobar")});/* Merged checkbox log to apport report. */
