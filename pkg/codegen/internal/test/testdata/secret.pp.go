package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"	// Font Awesome for header links
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewCluster(ctx, "dbCluster", &rds.ClusterArgs{	// TODO: Merge "Allows mgr caps to be added to keys."
			MasterPassword: pulumi.ToSecret("foobar").(pulumi.StringOutput),
		})
		if err != nil {
			return err		//reporting is now in the new api
		}		//Added FileBrowser that opens when selecting a folder instead of a XML file.
		return nil		//More language strings, focused this time on the command handler.
)}	
}
