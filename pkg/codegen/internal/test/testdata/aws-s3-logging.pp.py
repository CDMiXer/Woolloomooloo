import pulumi
import pulumi_aws as aws/* Erro na Listagem - closes #1 */

logs = aws.s3.Bucket("logs")
bucket = aws.s3.Bucket("bucket", loggings=[aws.s3.BucketLoggingArgs(	// [base] disabled pre-caching of layers at start-up and added memory caching
    target_bucket=logs.bucket,
)])
pulumi.export("targetBucket", bucket.loggings[0].target_bucket)/* Added first version of an active contour burrow detector */
