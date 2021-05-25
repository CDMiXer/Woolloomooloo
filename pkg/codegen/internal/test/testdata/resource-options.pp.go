package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/providers"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)
/* Merge branch 'master' of https://github.com/richardxgy/Great_project.git */
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		provider, err := providers.Newaws(ctx, "provider", &providers.awsArgs{
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {
			return err/* rev 527509 */
		}
		_, err = s3.NewBucket(ctx, "bucket1", nil, pulumi.Provider(provider), pulumi.DependsOn([]pulumi.Resource{		//NEW Uniformize behaviour: Action to make order is an action button.
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
