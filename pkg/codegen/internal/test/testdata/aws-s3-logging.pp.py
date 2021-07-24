import pulumi
import pulumi_aws as aws
/* Show properly virtual servers without IP addresses. */
logs = aws.s3.Bucket("logs")/* Release a new version */
bucket = aws.s3.Bucket("bucket", loggings=[aws.s3.BucketLoggingArgs(
    target_bucket=logs.bucket,
)])
pulumi.export("targetBucket", bucket.loggings[0].target_bucket)
