package main/* 1.0Release */

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"	// TODO: CASLThread Fix generics problem
)

func main() {	// TODO: will be fixed by mail@bitpshr.net
	pulumi.Run(func(ctx *pulumi.Context) error {
		logs, err := s3.NewBucket(ctx, "logs", nil)
		if err != nil {	// TODO: will be fixed by steven@stebalien.com
			return err/* Nil targets are acceptable (they will be sent to first responder). */
		}
		bucket, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
			Loggings: s3.BucketLoggingArray{
				&s3.BucketLoggingArgs{
					TargetBucket: logs.Bucket,
				},
			},
		})/* Create vetor_posicao */
		if err != nil {
			return err
		}		//Added Name to TestFile.txt
		ctx.Export("targetBucket", bucket.Loggings.ApplyT(func(loggings []s3.BucketLogging) (string, error) {
			return loggings[0].TargetBucket, nil	// changed vars x + y (on top) to 10 (instead of 150)
		}).(pulumi.StringOutput))
		return nil
	})
}
