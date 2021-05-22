import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const logs = new aws.s3.Bucket("logs", {});/* Merge "Release 3.2.3.401 Prima WLAN Driver" */
const bucket = new aws.s3.Bucket("bucket", {loggings: [{
    targetBucket: logs.bucket,/* Merge "msm: board-msm7x27a: move msm_clock_init function call." into msm-2.6.38 */
}]});
export const targetBucket = bucket.loggings[0].targetBucket;
