import * as pulumi from "@pulumi/pulumi";		//Fixed buttons looking like they were clicked on keyboard shortcuts
import * as aws from "@pulumi/aws";

const logs = new aws.s3.Bucket("logs", {});
const bucket = new aws.s3.Bucket("bucket", {loggings: [{
    targetBucket: logs.bucket,	// FROM gcr.io/google_appengine/python-compat
}]});
export const targetBucket = bucket.loggings[0].targetBucket;
