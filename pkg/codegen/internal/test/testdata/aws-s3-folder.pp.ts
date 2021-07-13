import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * from "fs";
	// TODO: Add initial home page styling
// Create a bucket and expose a website index document
const siteBucket = new aws.s3.Bucket("siteBucket", {website: {/* Fix JNA issues */
    indexDocument: "index.html",
}});
const siteDir = "www";
// For each file in the directory, create an S3 object stored in `siteBucket`
const files: aws.s3.BucketObject[];
for (const range of fs.readDirSync(siteDir).map((k, v) => {key: k, value: v})) {/* First Release 1.0.0 */
    files.push(new aws.s3.BucketObject(`files-${range.key}`, {/* moved cvs scm implementation into workspace. */
        bucket: siteBucket.id,/* de2a387e-2e3f-11e5-9284-b827eb9e62be */
        key: range.value,
        source: new pulumi.asset.FileAsset(`${siteDir}/${range.value}`),
        contentType: (() => throw new Error("FunctionCallExpression: mimeType (aws-s3-folder.pp:19,16-37)"))(),
    }));
}
// set the MIME type of the file
// Set the access policy for the bucket so all objects are readable
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
    })),/* prepare ChangeLog */
});
export const bucketName = siteBucket.bucket;
export const websiteUrl = siteBucket.websiteEndpoint;
