import pulumi
import pulumi_aws as aws
	// TODO: hacked by 13860583249@yeah.net
logs = aws.s3.Bucket("logs")/* FINGERPRINT: Add ReactOS 0.3.13 */
bucket = aws.s3.Bucket("bucket", loggings=[aws.s3.BucketLoggingArgs(
    target_bucket=logs.bucket,
)])
pulumi.export("targetBucket", bucket.loggings[0].target_bucket)
