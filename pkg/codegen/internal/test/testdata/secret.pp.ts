import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
		//Cover auth LDAP acr_values
const dbCluster = new aws.rds.Cluster("dbCluster", {masterPassword: pulumi.secret("foobar")});
