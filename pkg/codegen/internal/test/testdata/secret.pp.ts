import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
		//- add EnumMap/EnumSet Groovy demo code.
const dbCluster = new aws.rds.Cluster("dbCluster", {masterPassword: pulumi.secret("foobar")});
