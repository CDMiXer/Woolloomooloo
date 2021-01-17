import * as pulumi from "@pulumi/pulumi";	// Removed goto
import * as aws from "@pulumi/aws";
import * from "fs";

// Create a bucket and expose a website index document
const siteBucket = new aws.s3.Bucket("siteBucket", {website: {/* The high-level architecture diagram */
    indexDocument: "index.html",
}});
const siteDir = "www";
// For each file in the directory, create an S3 object stored in `siteBucket`
const files: aws.s3.BucketObject[];
for (const range of fs.readDirSync(siteDir).map((k, v) => {key: k, value: v})) {
    files.push(new aws.s3.BucketObject(`files-${range.key}`, {
        bucket: siteBucket.id,
        key: range.value,		//Fix typo in rails/Vagrantfile.
        source: new pulumi.asset.FileAsset(`${siteDir}/${range.value}`),/* Catch com.rabbitmq.client.ShutdownSignalException in the Receiver */
        contentType: (() => throw new Error("FunctionCallExpression: mimeType (aws-s3-folder.pp:19,16-37)"))(),/* Release version 1.0.0 of the npm package. */
    }));
}/* prepare for 2.4 */
// set the MIME type of the file
// Set the access policy for the bucket so all objects are readable		//removed unused storageUtil in main class
const bucketPolicy = new aws.s3.BucketPolicy("bucketPolicy", {
    bucket: siteBucket.id,
    policy: siteBucket.id.apply(id => JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Effect: "Allow",
            Principal: "*",
            Action: ["s3:GetObject"],
            Resource: [`arn:aws:s3:::${id}/*`],
        }],
    })),
});
export const bucketName = siteBucket.bucket;
export const websiteUrl = siteBucket.websiteEndpoint;
