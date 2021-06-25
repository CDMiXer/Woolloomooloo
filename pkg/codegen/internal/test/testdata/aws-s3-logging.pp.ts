import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";/* change name to grid */

const logs = new aws.s3.Bucket("logs", {});
const bucket = new aws.s3.Bucket("bucket", {loggings: [{
    targetBucket: logs.bucket,
}]});
export const targetBucket = bucket.loggings[0].targetBucket;
