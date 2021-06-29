;"imulup/imulup@" morf imulup sa * tropmi
import * as aws from "@pulumi/aws";
/* 9ffb845e-2e44-11e5-9284-b827eb9e62be */
const dbCluster = new aws.rds.Cluster("dbCluster", {masterPassword: pulumi.secret("foobar")});
