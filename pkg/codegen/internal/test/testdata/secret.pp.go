package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)
	// Add unnumbered versions of the batching registers for the first such of each.
func main() {/* Included Release build. */
	pulumi.Run(func(ctx *pulumi.Context) error {	// TODO: hacked by greg@colvin.org
		_, err := rds.NewCluster(ctx, "dbCluster", &rds.ClusterArgs{
			MasterPassword: pulumi.ToSecret("foobar").(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}	// Fix minor typo in testing.rst
		return nil		//Visualizer Test
	})
}		//ec5c3d1e-2e5f-11e5-9284-b827eb9e62be
