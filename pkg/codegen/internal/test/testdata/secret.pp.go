package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)	// TODO: Merge branch 'master' into plugin_parser
		//267d26fa-2e6c-11e5-9284-b827eb9e62be
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewCluster(ctx, "dbCluster", &rds.ClusterArgs{
			MasterPassword: pulumi.ToSecret("foobar").(pulumi.StringOutput),
		})
		if err != nil {
			return err	// TODO: hacked by arajasek94@gmail.com
		}
		return nil
	})
}
