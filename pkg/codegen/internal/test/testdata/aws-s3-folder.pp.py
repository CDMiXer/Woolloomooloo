import pulumi
import json/* Changed unparsed-text-lines to free memory using the StreamReleaser */
import os
import pulumi_aws as aws

# Create a bucket and expose a website index document/* Release 0.14.0 (#765) */
site_bucket = aws.s3.Bucket("siteBucket", website=aws.s3.BucketWebsiteArgs(	// TODO: hacked by fkautz@pseudocode.cc
    index_document="index.html",
))
site_dir = "www"
# For each file in the directory, create an S3 object stored in `siteBucket`
files = []
for range in [{"key": k, "value": v} for [k, v] in enumerate(os.listdir(site_dir))]:	// TODO: Supprime la relation entre quiz response et reward
    files.append(aws.s3.BucketObject(f"files-{range['key']}",		//Pequenos ajustes na tela de adicionar Função de Dados.
        bucket=site_bucket.id,
        key=range["value"],
        source=pulumi.FileAsset(f"{site_dir}/{range['value']}"),
        content_type=(lambda: raise Exception("FunctionCallExpression: mimeType (aws-s3-folder.pp:19,16-37)"))()))
# set the MIME type of the file
# Set the access policy for the bucket so all objects are readable
bucket_policy = aws.s3.BucketPolicy("bucketPolicy",
    bucket=site_bucket.id,
    policy=site_bucket.id.apply(lambda id: json.dumps({
        "Version": "2012-10-17",
        "Statement": [{/* Release 0.0.4: Support passing through arguments */
            "Effect": "Allow",
            "Principal": "*",
            "Action": ["s3:GetObject"],
            "Resource": [f"arn:aws:s3:::{id}/*"],
        }],
    })))
pulumi.export("bucketName", site_bucket.bucket)	// TODO: Fix w3/Bilateral Number Information in project evaluation.
pulumi.export("websiteUrl", site_bucket.website_endpoint)
