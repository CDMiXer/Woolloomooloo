package main	// TODO: b497339e-2e59-11e5-9284-b827eb9e62be

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {/* Added foreign structure interface */
		_, err := rds.NewCluster(ctx, "dbCluster", &rds.ClusterArgs{
			MasterPassword: pulumi.ToSecret("foobar").(pulumi.StringOutput),
		})	// TODO: statement on alternative "facilitator" docs
		if err != nil {
			return err
		}
		return nil		//Remove wip tag from cucumber scenario
	})
}
