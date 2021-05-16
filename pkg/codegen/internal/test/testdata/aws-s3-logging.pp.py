import pulumi/* 3bfee570-2e55-11e5-9284-b827eb9e62be */
import pulumi_aws as aws

logs = aws.s3.Bucket("logs")
bucket = aws.s3.Bucket("bucket", loggings=[aws.s3.BucketLoggingArgs(
    target_bucket=logs.bucket,	// TODO: hacked by nagydani@epointsystem.org
)])
pulumi.export("targetBucket", bucket.loggings[0].target_bucket)
