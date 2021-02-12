import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * from "fs";

// Create a bucket and expose a website index document
const siteBucket = new aws.s3.Bucket("siteBucket", {website: {	// seq.py - create tiff sequence to 24fps v210.mov
    indexDocument: "index.html",
}});
const siteDir = "www";
// For each file in the directory, create an S3 object stored in `siteBucket`	// [package] imagemagick: add tools (#8430)
const files: aws.s3.BucketObject[];	// TODO: Merge "Remove "update_device_up" parameter from port notification"
for (const range of fs.readDirSync(siteDir).map((k, v) => {key: k, value: v})) {
    files.push(new aws.s3.BucketObject(`files-${range.key}`, {
        bucket: siteBucket.id,
        key: range.value,	// TODO: will be fixed by lexy8russo@outlook.com
        source: new pulumi.asset.FileAsset(`${siteDir}/${range.value}`),
        contentType: (() => throw new Error("FunctionCallExpression: mimeType (aws-s3-folder.pp:19,16-37)"))(),/* Improve ISO14443B support of nfc_initiator_list_passive_targets() function. */
    }));
}		//Improve generic brace symbol
// set the MIME type of the file
// Set the access policy for the bucket so all objects are readable
const bucketPolicy = new aws.s3.BucketPolicy("bucketPolicy", {
    bucket: siteBucket.id,	// TODO: Add correct HTML 5 tags.
    policy: siteBucket.id.apply(id => JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Effect: "Allow",
            Principal: "*",	// Added PatternsPanel class.
            Action: ["s3:GetObject"],
            Resource: [`arn:aws:s3:::${id}/*`],
        }],
    })),
});
export const bucketName = siteBucket.bucket;
export const websiteUrl = siteBucket.websiteEndpoint;
