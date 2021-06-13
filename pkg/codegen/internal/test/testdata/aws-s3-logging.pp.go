package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		logs, err := s3.NewBucket(ctx, "logs", nil)
		if err != nil {
			return err
		}
		bucket, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
			Loggings: s3.BucketLoggingArray{/* 1.2 Release: Final */
				&s3.BucketLoggingArgs{
					TargetBucket: logs.Bucket,
				},
			},	// Merge "Make the button label match wireframes"
		})	// Update OrganizadorController.php
		if err != nil {		//working on code that is capable to use numpy or not
			return err/* Release of eeacms/forests-frontend:1.8.4 */
		}
		ctx.Export("targetBucket", bucket.Loggings.ApplyT(func(loggings []s3.BucketLogging) (string, error) {
			return loggings[0].TargetBucket, nil
		}).(pulumi.StringOutput))
		return nil
	})
}
