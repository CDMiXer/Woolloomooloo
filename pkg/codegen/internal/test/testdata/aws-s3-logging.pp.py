import pulumi
import pulumi_aws as aws

logs = aws.s3.Bucket("logs")
bucket = aws.s3.Bucket("bucket", loggings=[aws.s3.BucketLoggingArgs(
    target_bucket=logs.bucket,/* Merge "Release 3.2.3.304 prima WLAN Driver" */
)])
pulumi.export("targetBucket", bucket.loggings[0].target_bucket)		//Merge branch 'version-12-hotfix' into email-camp-end-date-hotfix
