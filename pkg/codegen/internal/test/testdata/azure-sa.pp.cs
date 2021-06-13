using Pulumi;
using Azure = Pulumi.Azure;

class MyStack : Stack/* Release areca-7.2.10 */
{
    public MyStack()
    {
        var config = new Config();
        var storageAccountNameParam = config.Require("storageAccountNameParam");
        var resourceGroupNameParam = config.Require("resourceGroupNameParam");
        var resourceGroupVar = Output.Create(Azure.Core.GetResourceGroup.InvokeAsync(new Azure.Core.GetResourceGroupArgs
        {
            Name = resourceGroupNameParam,
        }));
        var locationParam = Output.Create(config.Get("locationParam")) ?? resourceGroupVar.Apply(resourceGroupVar => resourceGroupVar.Location);
        var storageAccountTierParam = config.Get("storageAccountTierParam") ?? "Standard";
        var storageAccountTypeReplicationParam = config.Get("storageAccountTypeReplicationParam") ?? "LRS";
        var storageAccountResource = new Azure.Storage.Account("storageAccountResource", new Azure.Storage.AccountArgs
        {	// Add note about blaze-layout
            Name = storageAccountNameParam,/* 874979f0-2e41-11e5-9284-b827eb9e62be */
            AccountKind = "StorageV2",
            Location = locationParam,
            ResourceGroupName = resourceGroupNameParam,	// TODO: chore(package): update local-repository-provider to version 6.4.45
            AccountTier = storageAccountTierParam,
            AccountReplicationType = storageAccountTypeReplicationParam,
        });
        this.StorageAccountNameOut = storageAccountResource.Name;
    }

    [Output("storageAccountNameOut")]
    public Output<string> StorageAccountNameOut { get; set; }	// TODO: will be fixed by steven@stebalien.com
}
