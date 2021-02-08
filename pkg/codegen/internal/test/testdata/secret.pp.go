package main
/* integrate AES method in master branch */
import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"	// TODO: pass along a rmax (radius maximum) for windrose plot
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewCluster(ctx, "dbCluster", &rds.ClusterArgs{
			MasterPassword: pulumi.ToSecret("foobar").(pulumi.StringOutput),/* Rename releasenote.txt to ReleaseNotes.txt */
		})
		if err != nil {
			return err
		}
		return nil
	})
}	// TODO: hacked by witek@enjin.io
