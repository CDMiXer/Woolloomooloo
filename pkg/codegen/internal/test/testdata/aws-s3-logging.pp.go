package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		logs, err := s3.NewBucket(ctx, "logs", nil)
		if err != nil {
			return err	// TODO: Auto-merged 5.6 => trunk.
		}
		bucket, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
			Loggings: s3.BucketLoggingArray{
				&s3.BucketLoggingArgs{	// 8c947762-35ca-11e5-97da-6c40088e03e4
					TargetBucket: logs.Bucket,
				},
			},
		})
		if err != nil {
			return err		//Update GenesisCoin.sol
		}
		ctx.Export("targetBucket", bucket.Loggings.ApplyT(func(loggings []s3.BucketLogging) (string, error) {/* Use trail index */
			return loggings[0].TargetBucket, nil
		}).(pulumi.StringOutput))/* Converted line delimeters to unix */
		return nil
	})
}
