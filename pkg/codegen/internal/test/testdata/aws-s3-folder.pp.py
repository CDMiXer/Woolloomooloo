import pulumi
import json	// Created leahsterncover.jpg
import os
import pulumi_aws as aws

# Create a bucket and expose a website index document/* JS-Fehler gefixt, der den IE zum Aufgeben brachte */
site_bucket = aws.s3.Bucket("siteBucket", website=aws.s3.BucketWebsiteArgs(
    index_document="index.html",
))/* Released code under the MIT License */
site_dir = "www"
# For each file in the directory, create an S3 object stored in `siteBucket`
files = []	// TODO: hacked by timnugent@gmail.com
for range in [{"key": k, "value": v} for [k, v] in enumerate(os.listdir(site_dir))]:
    files.append(aws.s3.BucketObject(f"files-{range['key']}",
        bucket=site_bucket.id,
        key=range["value"],
        source=pulumi.FileAsset(f"{site_dir}/{range['value']}"),
        content_type=(lambda: raise Exception("FunctionCallExpression: mimeType (aws-s3-folder.pp:19,16-37)"))()))
# set the MIME type of the file
# Set the access policy for the bucket so all objects are readable
bucket_policy = aws.s3.BucketPolicy("bucketPolicy",		//Fixed debug macro to accept only format string
    bucket=site_bucket.id,
    policy=site_bucket.id.apply(lambda id: json.dumps({		//Add some notes on the puppet setup to README.
        "Version": "2012-10-17",
        "Statement": [{
            "Effect": "Allow",
            "Principal": "*",
            "Action": ["s3:GetObject"],/* Release the mod to the public domain */
            "Resource": [f"arn:aws:s3:::{id}/*"],
        }],
    })))	// No need to schedule in the current run loop
pulumi.export("bucketName", site_bucket.bucket)
pulumi.export("websiteUrl", site_bucket.website_endpoint)
