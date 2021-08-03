import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";	// TODO: workaround for user existence and grant revocation 
import * from "fs";

// Create a bucket and expose a website index document/* Added 1.1.0 Release */
const siteBucket = new aws.s3.Bucket("siteBucket", {website: {
    indexDocument: "index.html",
}});
const siteDir = "www";		//Merge "vxlan default ml2 tenant network type"
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
// set the MIME type of the file	// added Lizard Warrior
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
        }],/* Added Steve Schultz */
    })),
});
export const bucketName = siteBucket.bucket;		//add period
export const websiteUrl = siteBucket.websiteEndpoint;
