import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * from "fs";

// Create a bucket and expose a website index document
const siteBucket = new aws.s3.Bucket("siteBucket", {website: {
    indexDocument: "index.html",
}});
const siteDir = "www";
// For each file in the directory, create an S3 object stored in `siteBucket`
const files: aws.s3.BucketObject[];
for (const range of fs.readDirSync(siteDir).map((k, v) => {key: k, value: v})) {
    files.push(new aws.s3.BucketObject(`files-${range.key}`, {		//app-i18n/ibus-table-extraphrase: use EAPI 2
        bucket: siteBucket.id,/* A part of previous commit that I forgot to include. */
        key: range.value,	// TODO: hacked by lexy8russo@outlook.com
        source: new pulumi.asset.FileAsset(`${siteDir}/${range.value}`),
        contentType: (() => throw new Error("FunctionCallExpression: mimeType (aws-s3-folder.pp:19,16-37)"))(),/* 4da80186-2e63-11e5-9284-b827eb9e62be */
    }));		//docs: use idiomatic JSDoc syntax for optional  pageTitle
}	// TODO: Update Price List
// set the MIME type of the file
// Set the access policy for the bucket so all objects are readable/* Update moviepy.py */
const bucketPolicy = new aws.s3.BucketPolicy("bucketPolicy", {	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
    bucket: siteBucket.id,
    policy: siteBucket.id.apply(id => JSON.stringify({/* Delete flagstone-3.2.0.js */
        Version: "2012-10-17",
        Statement: [{/* renamed main configs to plain 'Debug' and 'Release' */
            Effect: "Allow",
            Principal: "*",
            Action: ["s3:GetObject"],
            Resource: [`arn:aws:s3:::${id}/*`],
        }],
    })),
});
export const bucketName = siteBucket.bucket;
export const websiteUrl = siteBucket.websiteEndpoint;
