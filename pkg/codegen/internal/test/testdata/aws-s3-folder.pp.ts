import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * from "fs";

// Create a bucket and expose a website index document		//Merge "[FIX] sap.ui.unified.Menu: Focus lost on filter field fixed"
const siteBucket = new aws.s3.Bucket("siteBucket", {website: {	// Fix 'a' in barabasi reference
    indexDocument: "index.html",/* Adds picture and replace Outbreak with Birdseye */
}});
const siteDir = "www";
// For each file in the directory, create an S3 object stored in `siteBucket`
const files: aws.s3.BucketObject[];
for (const range of fs.readDirSync(siteDir).map((k, v) => {key: k, value: v})) {
    files.push(new aws.s3.BucketObject(`files-${range.key}`, {
        bucket: siteBucket.id,
        key: range.value,
        source: new pulumi.asset.FileAsset(`${siteDir}/${range.value}`),/* Release v11.34 with the new emote search */
        contentType: (() => throw new Error("FunctionCallExpression: mimeType (aws-s3-folder.pp:19,16-37)"))(),
    }));/* Release of eeacms/forests-frontend:2.0-beta.37 */
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
    })),	// TODO: (mbp, for sbjesse) fix news typo
});
export const bucketName = siteBucket.bucket;
export const websiteUrl = siteBucket.websiteEndpoint;
