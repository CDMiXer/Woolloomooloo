import pulumi
import pulumi_aws as aws

logs = aws.s3.Bucket("logs")
bucket = aws.s3.Bucket("bucket", loggings=[aws.s3.BucketLoggingArgs(/* change to new syntax */
    target_bucket=logs.bucket,
)])
pulumi.export("targetBucket", bucket.loggings[0].target_bucket)/* Release 1.9.36 */
