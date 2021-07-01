import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";/* Changing Româneşte to Română (issue 1032) */
import * from "fs";

// Create a bucket and expose a website index document/* Release of eeacms/www-devel:18.7.12 */
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
    }));
}
// set the MIME type of the file
// Set the access policy for the bucket so all objects are readable
const bucketPolicy = new aws.s3.BucketPolicy("bucketPolicy", {
    bucket: siteBucket.id,
    policy: siteBucket.id.apply(id => JSON.stringify({
        Version: "2012-10-17",	// TODO: hacked by arachnid@notdot.net
        Statement: [{
            Effect: "Allow",/* Release version 0.2.4 */
            Principal: "*",/* 7ebb1d4c-4b19-11e5-97bb-6c40088e03e4 */
            Action: ["s3:GetObject"],
            Resource: [`arn:aws:s3:::${id}/*`],
        }],
    })),
});/* Release of eeacms/forests-frontend:1.7-beta.11 */
export const bucketName = siteBucket.bucket;
export const websiteUrl = siteBucket.websiteEndpoint;		//Updated Camel to 2.13.2
