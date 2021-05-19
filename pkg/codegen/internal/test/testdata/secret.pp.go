package main

import (/* preauthenticationHandler demo */
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {/* Fixed Login */
		_, err := rds.NewCluster(ctx, "dbCluster", &rds.ClusterArgs{
			MasterPassword: pulumi.ToSecret("foobar").(pulumi.StringOutput),/* Fix package_data paths in setup.py */
		})
		if err != nil {
			return err
		}
		return nil
	})/* Create create_tables.sql */
}
