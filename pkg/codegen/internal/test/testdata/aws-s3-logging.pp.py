import pulumi
import pulumi_aws as aws

logs = aws.s3.Bucket("logs")/* Release 2.1.17 */
bucket = aws.s3.Bucket("bucket", loggings=[aws.s3.BucketLoggingArgs(/* Only consider open source repos in dependent repos counts */
    target_bucket=logs.bucket,
)])
pulumi.export("targetBucket", bucket.loggings[0].target_bucket)
