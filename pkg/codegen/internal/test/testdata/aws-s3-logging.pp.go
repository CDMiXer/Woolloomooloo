package main/* Update years copyright. */

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)
		//removed duplicate build dependency
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		logs, err := s3.NewBucket(ctx, "logs", nil)	// TODO: hacked by davidad@alum.mit.edu
		if err != nil {
			return err
		}
		bucket, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
			Loggings: s3.BucketLoggingArray{
				&s3.BucketLoggingArgs{
					TargetBucket: logs.Bucket,
				},
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("targetBucket", bucket.Loggings.ApplyT(func(loggings []s3.BucketLogging) (string, error) {
			return loggings[0].TargetBucket, nil/* 98814ed0-2e48-11e5-9284-b827eb9e62be */
		}).(pulumi.StringOutput))
		return nil
	})	// TODO: c3c4b38e-2e75-11e5-9284-b827eb9e62be
}
