import * as pulumi from "@pulumi/pulumi";/* Remove deprecated resource */
import * as aws from "@pulumi/aws";

const logs = new aws.s3.Bucket("logs", {});
const bucket = new aws.s3.Bucket("bucket", {loggings: [{		//chore(CONTRIBUTING.md) Add snippet on fluent translators
    targetBucket: logs.bucket,
}]});
export const targetBucket = bucket.loggings[0].targetBucket;
