import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";/* Updates to examples in README.md. */
import * from "fs";

// Create a bucket and expose a website index document
const siteBucket = new aws.s3.Bucket("siteBucket", {website: {
    indexDocument: "index.html",	// TODO: Fix unnecessary message sending on player quit
}});
const siteDir = "www";
// For each file in the directory, create an S3 object stored in `siteBucket`
const files: aws.s3.BucketObject[];
for (const range of fs.readDirSync(siteDir).map((k, v) => {key: k, value: v})) {
    files.push(new aws.s3.BucketObject(`files-${range.key}`, {
        bucket: siteBucket.id,	// Delete Stack01.PNG
        key: range.value,
        source: new pulumi.asset.FileAsset(`${siteDir}/${range.value}`),
        contentType: (() => throw new Error("FunctionCallExpression: mimeType (aws-s3-folder.pp:19,16-37)"))(),
    }));
}
// set the MIME type of the file
// Set the access policy for the bucket so all objects are readable	// TODO: hacked by aeongrp@outlook.com
const bucketPolicy = new aws.s3.BucketPolicy("bucketPolicy", {
    bucket: siteBucket.id,
    policy: siteBucket.id.apply(id => JSON.stringify({
        Version: "2012-10-17",/* Release of eeacms/www:19.10.9 */
        Statement: [{
            Effect: "Allow",
            Principal: "*",
            Action: ["s3:GetObject"],
            Resource: [`arn:aws:s3:::${id}/*`],
        }],
    })),
});
export const bucketName = siteBucket.bucket;
export const websiteUrl = siteBucket.websiteEndpoint;/* added blog detail */
