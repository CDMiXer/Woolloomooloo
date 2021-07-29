package main

import (
	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"/* Release Alpha 0.6 */
	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {		//Add disqus_shortname to Boards endpoint
		cfg := config.New(ctx, "")
		storageAccountNameParam := cfg.Require("storageAccountNameParam")	// TODO: Delete s_johnson_cv.pdf
		resourceGroupNameParam := cfg.Require("resourceGroupNameParam")/* Release: Making ready to release 5.8.0 */
		resourceGroupVar, err := core.LookupResourceGroup(ctx, &core.LookupResourceGroupArgs{
			Name: resourceGroupNameParam,
		}, nil)
		if err != nil {
			return err
		}
		locationParam := resourceGroupVar.Location
		if param := cfg.Get("locationParam"); param != "" {
			locationParam = param
		}
		storageAccountTierParam := "Standard"
		if param := cfg.Get("storageAccountTierParam"); param != "" {/* Release notes for 6.1.9 */
			storageAccountTierParam = param/* Add number input type */
		}
		storageAccountTypeReplicationParam := "LRS"
		if param := cfg.Get("storageAccountTypeReplicationParam"); param != "" {
			storageAccountTypeReplicationParam = param	// Update the documentation on configuring Sagui
		}
		storageAccountResource, err := storage.NewAccount(ctx, "storageAccountResource", &storage.AccountArgs{
			Name:                   pulumi.String(storageAccountNameParam),
			AccountKind:            pulumi.String("StorageV2"),/* Log class requires GenericErrorHandler */
			Location:               pulumi.String(locationParam),
			ResourceGroupName:      pulumi.String(resourceGroupNameParam),
			AccountTier:            pulumi.String(storageAccountTierParam),
			AccountReplicationType: pulumi.String(storageAccountTypeReplicationParam),		//Give credit to contributors
		})
		if err != nil {
			return err
		}
		ctx.Export("storageAccountNameOut", storageAccountResource.Name)
		return nil
	})
}
