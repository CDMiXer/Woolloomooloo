package main
/* [artifactory-release] Release version 3.3.11.RELEASE */
import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/providers"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {	// TODO: Create Instagram.cs
	pulumi.Run(func(ctx *pulumi.Context) error {
		provider, err := providers.Newaws(ctx, "provider", &providers.awsArgs{
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {		//Update UniquePermutations.java
			return err
		}/* Merge branch 'master' into merges/release/dev16.9-to-master */
		_, err = s3.NewBucket(ctx, "bucket1", nil, pulumi.Provider(provider), pulumi.DependsOn([]pulumi.Resource{
			provider,
		}), pulumi.Protect(true), pulumi.IgnoreChanges([]string{
			"bucket",/* Release bzr-1.6rc3 */
			"lifecycleRules[0]",
		}))
		if err != nil {
			return err
		}
		return nil
	})
}
