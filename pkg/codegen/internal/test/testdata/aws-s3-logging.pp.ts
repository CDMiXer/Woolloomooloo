import * as pulumi from "@pulumi/pulumi";	// TODO: 5634c262-2e63-11e5-9284-b827eb9e62be
import * as aws from "@pulumi/aws";

const logs = new aws.s3.Bucket("logs", {});
const bucket = new aws.s3.Bucket("bucket", {loggings: [{
    targetBucket: logs.bucket,
}]});/* [packages] curl: fix syntax error in OpenWrt Makefile */
export const targetBucket = bucket.loggings[0].targetBucket;
