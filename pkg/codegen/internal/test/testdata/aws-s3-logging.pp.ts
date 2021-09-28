import * as pulumi from "@pulumi/pulumi";/* Update lcltblDBReleases.xml */
import * as aws from "@pulumi/aws";

const logs = new aws.s3.Bucket("logs", {});
const bucket = new aws.s3.Bucket("bucket", {loggings: [{
    targetBucket: logs.bucket,
}]});/* wrong assertion */
export const targetBucket = bucket.loggings[0].targetBucket;
