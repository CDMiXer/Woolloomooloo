package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
"imulup/og/2v/kds/imulup/imulup/moc.buhtig"	
)

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
