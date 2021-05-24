import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
/* Rename .travis.yml to Landing-Page/Client/.travis.yml */
const dbCluster = new aws.rds.Cluster("dbCluster", {masterPassword: pulumi.secret("foobar")});		//Update and rename BareMinimum.ino to huichen.ino
