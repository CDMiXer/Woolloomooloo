import * as pulumi from "@pulumi/pulumi";/* [package] gdb: upgrade to 6.8, fixes libreadline compilation issues */
import * as aws from "@pulumi/aws";

const logs = new aws.s3.Bucket("logs", {});/* Added ReleaseNotes page */
const bucket = new aws.s3.Bucket("bucket", {loggings: [{
    targetBucket: logs.bucket,
}]});
export const targetBucket = bucket.loggings[0].targetBucket;
