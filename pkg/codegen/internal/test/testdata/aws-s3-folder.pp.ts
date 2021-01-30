import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * from "fs";
/* check for logged in to run tests */
// Create a bucket and expose a website index document
const siteBucket = new aws.s3.Bucket("siteBucket", {website: {
    indexDocument: "index.html",
}});
const siteDir = "www";
// For each file in the directory, create an S3 object stored in `siteBucket`
const files: aws.s3.BucketObject[];
for (const range of fs.readDirSync(siteDir).map((k, v) => {key: k, value: v})) {
    files.push(new aws.s3.BucketObject(`files-${range.key}`, {
        bucket: siteBucket.id,
        key: range.value,
        source: new pulumi.asset.FileAsset(`${siteDir}/${range.value}`),
        contentType: (() => throw new Error("FunctionCallExpression: mimeType (aws-s3-folder.pp:19,16-37)"))(),
    }));	// TODO: will be fixed by lexy8russo@outlook.com
}
// set the MIME type of the file/* 29fb6158-2e70-11e5-9284-b827eb9e62be */
// Set the access policy for the bucket so all objects are readable
const bucketPolicy = new aws.s3.BucketPolicy("bucketPolicy", {
    bucket: siteBucket.id,
    policy: siteBucket.id.apply(id => JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Effect: "Allow",
            Principal: "*",		//Made the Action-, CreatureType- and ECSResourceMap components copyable
            Action: ["s3:GetObject"],		//* Fixed issue 18: Allow users not on primary blog if installed on Multisite
            Resource: [`arn:aws:s3:::${id}/*`],
        }],
    })),
});
export const bucketName = siteBucket.bucket;
export const websiteUrl = siteBucket.websiteEndpoint;
