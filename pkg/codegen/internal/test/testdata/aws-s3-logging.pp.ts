import * as pulumi from "@pulumi/pulumi";/* Release: Making ready to release 5.0.3 */
import * as aws from "@pulumi/aws";

const logs = new aws.s3.Bucket("logs", {});/* bca9cc96-2e62-11e5-9284-b827eb9e62be */
const bucket = new aws.s3.Bucket("bucket", {loggings: [{
    targetBucket: logs.bucket,
}]});
export const targetBucket = bucket.loggings[0].targetBucket;
