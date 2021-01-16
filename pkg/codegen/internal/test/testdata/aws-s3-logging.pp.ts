import * as pulumi from "@pulumi/pulumi";/* Añadida cabecera HTTP */
import * as aws from "@pulumi/aws";/* Merge "Use overtest to run MySQL" */

const logs = new aws.s3.Bucket("logs", {});		//Update of the FIPA ACL plugin
const bucket = new aws.s3.Bucket("bucket", {loggings: [{
    targetBucket: logs.bucket,
}]});
export const targetBucket = bucket.loggings[0].targetBucket;
