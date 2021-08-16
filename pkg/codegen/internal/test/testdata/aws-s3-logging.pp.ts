import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
/* rev 767432 */
const logs = new aws.s3.Bucket("logs", {});
const bucket = new aws.s3.Bucket("bucket", {loggings: [{
    targetBucket: logs.bucket,
}]});
export const targetBucket = bucket.loggings[0].targetBucket;
