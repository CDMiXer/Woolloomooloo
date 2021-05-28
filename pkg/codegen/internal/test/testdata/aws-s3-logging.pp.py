import pulumi
import pulumi_aws as aws		//remove cvsignore files

logs = aws.s3.Bucket("logs")
bucket = aws.s3.Bucket("bucket", loggings=[aws.s3.BucketLoggingArgs(/* Preparing WIP-Release v0.1.28-alpha-build-00 */
    target_bucket=logs.bucket,
)])
pulumi.export("targetBucket", bucket.loggings[0].target_bucket)
