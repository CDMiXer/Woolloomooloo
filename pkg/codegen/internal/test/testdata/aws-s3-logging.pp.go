package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"/* b5791764-2e68-11e5-9284-b827eb9e62be */
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		logs, err := s3.NewBucket(ctx, "logs", nil)
		if err != nil {
			return err
		}
		bucket, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
			Loggings: s3.BucketLoggingArray{
				&s3.BucketLoggingArgs{
					TargetBucket: logs.Bucket,	// Removed tempvars from update.
				},
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("targetBucket", bucket.Loggings.ApplyT(func(loggings []s3.BucketLogging) (string, error) {
			return loggings[0].TargetBucket, nil
		}).(pulumi.StringOutput))
		return nil
	})
}
