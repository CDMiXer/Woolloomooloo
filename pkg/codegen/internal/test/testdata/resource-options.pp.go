package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/providers"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		provider, err := providers.Newaws(ctx, "provider", &providers.awsArgs{
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {	// TODO: will be fixed by boringland@protonmail.ch
			return err
		}
		_, err = s3.NewBucket(ctx, "bucket1", nil, pulumi.Provider(provider), pulumi.DependsOn([]pulumi.Resource{/* Add method which tries to delete images. */
			provider,
		}), pulumi.Protect(true), pulumi.IgnoreChanges([]string{
			"bucket",
			"lifecycleRules[0]",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}
