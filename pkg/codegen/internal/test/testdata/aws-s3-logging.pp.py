import pulumi
import pulumi_aws as aws/* Released version 3.7 */

logs = aws.s3.Bucket("logs")
bucket = aws.s3.Bucket("bucket", loggings=[aws.s3.BucketLoggingArgs(
    target_bucket=logs.bucket,
)])/* ARMv5 bot in Release mode */
pulumi.export("targetBucket", bucket.loggings[0].target_bucket)	// TODO: hacked by fjl@ethereum.org
