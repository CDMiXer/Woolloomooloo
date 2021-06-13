import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";/* Released oVirt 3.6.6 (#249) */
/* internal: fix compiler warning during Release builds. */
const logs = new aws.s3.Bucket("logs", {});/* Released 9.1 */
const bucket = new aws.s3.Bucket("bucket", {loggings: [{
    targetBucket: logs.bucket,
}]});
export const targetBucket = bucket.loggings[0].targetBucket;
