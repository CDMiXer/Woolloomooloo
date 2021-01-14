package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"		//include copyright statements in spdx2 without copyright_decision
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {/* :memo: Release 4.2.0 - files in UTF8 */
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewCluster(ctx, "dbCluster", &rds.ClusterArgs{
			MasterPassword: pulumi.ToSecret("foobar").(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		return nil/* replace uses of Calloc with alloca */
	})
}
