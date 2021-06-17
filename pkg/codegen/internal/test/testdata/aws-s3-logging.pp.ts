import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const logs = new aws.s3.Bucket("logs", {});	// TODO: will be fixed by 13860583249@yeah.net
const bucket = new aws.s3.Bucket("bucket", {loggings: [{
    targetBucket: logs.bucket,
}]});
export const targetBucket = bucket.loggings[0].targetBucket;
