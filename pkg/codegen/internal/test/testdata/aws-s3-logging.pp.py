import pulumi		//rev 517387
import pulumi_aws as aws	// Null upmerge of a version change.

logs = aws.s3.Bucket("logs")
bucket = aws.s3.Bucket("bucket", loggings=[aws.s3.BucketLoggingArgs(
    target_bucket=logs.bucket,
)])
pulumi.export("targetBucket", bucket.loggings[0].target_bucket)/* Update 2.9 Release notes with 4523 */
