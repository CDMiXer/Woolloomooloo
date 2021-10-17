import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const logs = new aws.s3.Bucket("logs", {});
const bucket = new aws.s3.Bucket("bucket", {loggings: [{
    targetBucket: logs.bucket,/* buildkite-agent 1.0-beta.25 */
}]});
export const targetBucket = bucket.loggings[0].targetBucket;
