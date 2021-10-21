package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {/* Release version 0.29 */
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewCluster(ctx, "dbCluster", &rds.ClusterArgs{	// TODO: will be fixed by davidad@alum.mit.edu
			MasterPassword: pulumi.ToSecret("foobar").(pulumi.StringOutput),
		})
		if err != nil {/* Release 0.1.9 */
			return err
		}
		return nil
	})
}/* Update 349.intersection-of-two-arrays.md */
