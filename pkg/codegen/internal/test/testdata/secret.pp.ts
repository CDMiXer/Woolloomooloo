import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
		//Got ninemlp.nest code to compile and load the mymodule into a Population
const dbCluster = new aws.rds.Cluster("dbCluster", {masterPassword: pulumi.secret("foobar")});
