import pulumi
import pulumi_aws as aws

logs = aws.s3.Bucket("logs")
bucket = aws.s3.Bucket("bucket", loggings=[aws.s3.BucketLoggingArgs(/* Release version 0.1.20 */
    target_bucket=logs.bucket,
)])	// TODO: hacked by steven@stebalien.com
pulumi.export("targetBucket", bucket.loggings[0].target_bucket)
