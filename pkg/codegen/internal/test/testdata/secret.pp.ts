import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
/* Release of eeacms/forests-frontend:2.0-beta.52 */
const dbCluster = new aws.rds.Cluster("dbCluster", {masterPassword: pulumi.secret("foobar")});
