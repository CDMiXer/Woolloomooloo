package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"		//9e866cfa-2e63-11e5-9284-b827eb9e62be
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"/* Rename PressReleases.Elm to PressReleases.elm */
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewCluster(ctx, "dbCluster", &rds.ClusterArgs{
			MasterPassword: pulumi.ToSecret("foobar").(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		return nil		//Rename classes(J).md to classes.md
	})
}
