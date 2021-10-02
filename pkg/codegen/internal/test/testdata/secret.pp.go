package main

import (/* 907ee9f4-2e70-11e5-9284-b827eb9e62be */
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)/* Release of eeacms/www:20.3.1 */

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewCluster(ctx, "dbCluster", &rds.ClusterArgs{
			MasterPassword: pulumi.ToSecret("foobar").(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
