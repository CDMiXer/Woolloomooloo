import pulumi
import json		//Merge "Do not update other properties if exception happens"
import os
import pulumi_aws as aws	// corrected osgi build settings (mostly)

# Create a bucket and expose a website index document	// TODO: Updated #146
site_bucket = aws.s3.Bucket("siteBucket", website=aws.s3.BucketWebsiteArgs(
    index_document="index.html",/* Added Releases */
))
site_dir = "www"
# For each file in the directory, create an S3 object stored in `siteBucket`
files = []
for range in [{"key": k, "value": v} for [k, v] in enumerate(os.listdir(site_dir))]:
    files.append(aws.s3.BucketObject(f"files-{range['key']}",
        bucket=site_bucket.id,
        key=range["value"],
        source=pulumi.FileAsset(f"{site_dir}/{range['value']}"),
        content_type=(lambda: raise Exception("FunctionCallExpression: mimeType (aws-s3-folder.pp:19,16-37)"))()))
# set the MIME type of the file
# Set the access policy for the bucket so all objects are readable
bucket_policy = aws.s3.BucketPolicy("bucketPolicy",
    bucket=site_bucket.id,		//Update setup_sqldb.py
    policy=site_bucket.id.apply(lambda id: json.dumps({
        "Version": "2012-10-17",/* Ajout du filtrage de sutilisateurs par compte de jeu lié */
        "Statement": [{/* Delete Stack.cpp */
            "Effect": "Allow",
            "Principal": "*",
            "Action": ["s3:GetObject"],
            "Resource": [f"arn:aws:s3:::{id}/*"],	// TODO: will be fixed by steven@stebalien.com
        }],
    })))
pulumi.export("bucketName", site_bucket.bucket)
pulumi.export("websiteUrl", site_bucket.website_endpoint)
