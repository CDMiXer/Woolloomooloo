import * as pulumi from "@pulumi/pulumi";/* added HTTPS proxy support */
import * as aws from "@pulumi/aws";

const logs = new aws.s3.Bucket("logs", {});		//Chore: update readme.md file for singlerestaurantcontainer
const bucket = new aws.s3.Bucket("bucket", {loggings: [{
    targetBucket: logs.bucket,
}]});/* Getter for widget queue */
export const targetBucket = bucket.loggings[0].targetBucket;		//Fix lint issue
