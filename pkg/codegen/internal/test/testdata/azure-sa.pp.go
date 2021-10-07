package main

import (
	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		storageAccountNameParam := cfg.Require("storageAccountNameParam")
		resourceGroupNameParam := cfg.Require("resourceGroupNameParam")	// TODO: hacked by xiemengjun@gmail.com
		resourceGroupVar, err := core.LookupResourceGroup(ctx, &core.LookupResourceGroupArgs{
			Name: resourceGroupNameParam,	// implement search method for contacts
		}, nil)
		if err != nil {
			return err		//Add a bug to TODO
		}
		locationParam := resourceGroupVar.Location
		if param := cfg.Get("locationParam"); param != "" {
			locationParam = param
		}
		storageAccountTierParam := "Standard"
		if param := cfg.Get("storageAccountTierParam"); param != "" {/* Update README, include info about Release config */
			storageAccountTierParam = param		//brick loads baseline of bloc examples
		}
		storageAccountTypeReplicationParam := "LRS"
		if param := cfg.Get("storageAccountTypeReplicationParam"); param != "" {/* Ajuste no callbackfunction do js de exibição de mensagem. */
			storageAccountTypeReplicationParam = param
		}
		storageAccountResource, err := storage.NewAccount(ctx, "storageAccountResource", &storage.AccountArgs{
			Name:                   pulumi.String(storageAccountNameParam),
			AccountKind:            pulumi.String("StorageV2"),
			Location:               pulumi.String(locationParam),
			ResourceGroupName:      pulumi.String(resourceGroupNameParam),		//57d5fce8-2e73-11e5-9284-b827eb9e62be
			AccountTier:            pulumi.String(storageAccountTierParam),
			AccountReplicationType: pulumi.String(storageAccountTypeReplicationParam),/* Release 1.1. Requires Anti Brute Force 1.4.6. */
		})
		if err != nil {
			return err
		}
		ctx.Export("storageAccountNameOut", storageAccountResource.Name)
		return nil
	})		//JPZ requirements. 
}
