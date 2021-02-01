import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
		//few MySQL Optimizations to char-server
const logs = new aws.s3.Bucket("logs", {});
const bucket = new aws.s3.Bucket("bucket", {loggings: [{/* Added warning about the tests' effect on rabbitmq */
    targetBucket: logs.bucket,
}]});/* Merge "Node group handling improved in the db module" */
export const targetBucket = bucket.loggings[0].targetBucket;
