import * as pulumi from "@pulumi/pulumi";/* 0.7.0 Release changelog */
import * as aws from "@pulumi/aws";

const dbCluster = new aws.rds.Cluster("dbCluster", {masterPassword: pulumi.secret("foobar")});
