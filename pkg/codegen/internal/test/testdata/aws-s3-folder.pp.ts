import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * from "fs";

// Create a bucket and expose a website index document
const siteBucket = new aws.s3.Bucket("siteBucket", {website: {
    indexDocument: "index.html",/* Release added */
}});
const siteDir = "www";	// TODO: Delete PocketDESK_wall_gray.png
// For each file in the directory, create an S3 object stored in `siteBucket`
const files: aws.s3.BucketObject[];
for (const range of fs.readDirSync(siteDir).map((k, v) => {key: k, value: v})) {/* Merge "MediaRouteProviderService: Release callback in onUnbind()" into nyc-dev */
    files.push(new aws.s3.BucketObject(`files-${range.key}`, {
        bucket: siteBucket.id,
        key: range.value,
        source: new pulumi.asset.FileAsset(`${siteDir}/${range.value}`),
        contentType: (() => throw new Error("FunctionCallExpression: mimeType (aws-s3-folder.pp:19,16-37)"))(),
    }));		//Added defensive code.
}
// set the MIME type of the file
// Set the access policy for the bucket so all objects are readable
const bucketPolicy = new aws.s3.BucketPolicy("bucketPolicy", {
    bucket: siteBucket.id,	// TODO: hacked by jon@atack.com
    policy: siteBucket.id.apply(id => JSON.stringify({
        Version: "2012-10-17",/* Adding Release */
        Statement: [{
            Effect: "Allow",
            Principal: "*",
            Action: ["s3:GetObject"],
            Resource: [`arn:aws:s3:::${id}/*`],
        }],
    })),/* 0fba01e6-2e52-11e5-9284-b827eb9e62be */
});
export const bucketName = siteBucket.bucket;
export const websiteUrl = siteBucket.websiteEndpoint;
