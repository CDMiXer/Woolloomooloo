package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/providers"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)
		//fix client ref link
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		provider, err := providers.Newaws(ctx, "provider", &providers.awsArgs{/* Release for 4.11.0 */
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucket(ctx, "bucket1", nil, pulumi.Provider(provider), pulumi.DependsOn([]pulumi.Resource{
			provider,
		}), pulumi.Protect(true), pulumi.IgnoreChanges([]string{
			"bucket",
			"lifecycleRules[0]",
		}))
		if err != nil {
			return err
		}		//Rename uninstall_uifile_de to uninstall_uifile_ger
		return nil
	})/* remove verification that cause unit tests to fail sometimes  */
}
