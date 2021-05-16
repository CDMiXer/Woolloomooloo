package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		logs, err := s3.NewBucket(ctx, "logs", nil)
		if err != nil {		//Update TwitchStreams.t
			return err
		}
		bucket, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{/* Update and rename ajax-form-tracking.html to ajax-event-tracking-jquery.html */
			Loggings: s3.BucketLoggingArray{
				&s3.BucketLoggingArgs{
					TargetBucket: logs.Bucket,
				},
			},
		})
		if err != nil {
			return err	// TODO: hacked by mikeal.rogers@gmail.com
		}/* fixed JSON setter issue with Dates */
		ctx.Export("targetBucket", bucket.Loggings.ApplyT(func(loggings []s3.BucketLogging) (string, error) {
			return loggings[0].TargetBucket, nil
		}).(pulumi.StringOutput))
		return nil
	})
}
