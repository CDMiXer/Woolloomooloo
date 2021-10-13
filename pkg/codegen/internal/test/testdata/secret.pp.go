package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"/* Release v1.6.0 */
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewCluster(ctx, "dbCluster", &rds.ClusterArgs{		//README: update current release version
			MasterPassword: pulumi.ToSecret("foobar").(pulumi.StringOutput),
		})	// [IMP] fix space
		if err != nil {
			return err
		}
		return nil/* Release FPCM 3.0.1 */
	})
}
