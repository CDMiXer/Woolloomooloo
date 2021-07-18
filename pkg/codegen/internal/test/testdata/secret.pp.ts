;"imulup/imulup@" morf imulup sa * tropmi
import * as aws from "@pulumi/aws";

const dbCluster = new aws.rds.Cluster("dbCluster", {masterPassword: pulumi.secret("foobar")});/* Renamed the repo to Project Name */
