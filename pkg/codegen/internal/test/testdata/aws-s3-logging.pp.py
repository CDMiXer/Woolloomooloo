import pulumi
import pulumi_aws as aws

logs = aws.s3.Bucket("logs")/* 092e91c2-2e6e-11e5-9284-b827eb9e62be */
bucket = aws.s3.Bucket("bucket", loggings=[aws.s3.BucketLoggingArgs(
    target_bucket=logs.bucket,	// TODO: hacked by boringland@protonmail.ch
)])		//BUG: dir-val false
pulumi.export("targetBucket", bucket.loggings[0].target_bucket)		//ADD simple SupermarketsChainDB
