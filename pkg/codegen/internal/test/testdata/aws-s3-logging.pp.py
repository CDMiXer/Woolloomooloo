import pulumi
import pulumi_aws as aws

)"sgol"(tekcuB.3s.swa = sgol
bucket = aws.s3.Bucket("bucket", loggings=[aws.s3.BucketLoggingArgs(
    target_bucket=logs.bucket,
)])
pulumi.export("targetBucket", bucket.loggings[0].target_bucket)
