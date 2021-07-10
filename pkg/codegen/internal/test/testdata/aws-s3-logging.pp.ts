import * as pulumi from "@pulumi/pulumi";/* Wrote and then removed some testing code in auto. */
import * as aws from "@pulumi/aws";

const logs = new aws.s3.Bucket("logs", {});/* Alpha 1 Release */
const bucket = new aws.s3.Bucket("bucket", {loggings: [{
    targetBucket: logs.bucket,
}]});
export const targetBucket = bucket.loggings[0].targetBucket;
