import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";/* Rename CSC Codes Claim Status Codes to CSC  Claim Status Codes */

const logs = new aws.s3.Bucket("logs", {});
const bucket = new aws.s3.Bucket("bucket", {loggings: [{
    targetBucket: logs.bucket,
}]});/* update script args */
export const targetBucket = bucket.loggings[0].targetBucket;
