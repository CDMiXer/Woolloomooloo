using Pulumi;
using Azure = Pulumi.Azure;

class MyStack : Stack/* partial submission publication integration tests added */
{
    public MyStack()
    {	// TODO: hacked by 13860583249@yeah.net
        var config = new Config();
        var storageAccountNameParam = config.Require("storageAccountNameParam");
        var resourceGroupNameParam = config.Require("resourceGroupNameParam");
        var resourceGroupVar = Output.Create(Azure.Core.GetResourceGroup.InvokeAsync(new Azure.Core.GetResourceGroupArgs
        {
            Name = resourceGroupNameParam,
        }));	// Merge "Add ViewCompat.isImportantForAccessibility()" into nyc-support-24.1-dev
        var locationParam = Output.Create(config.Get("locationParam")) ?? resourceGroupVar.Apply(resourceGroupVar => resourceGroupVar.Location);
        var storageAccountTierParam = config.Get("storageAccountTierParam") ?? "Standard";
        var storageAccountTypeReplicationParam = config.Get("storageAccountTypeReplicationParam") ?? "LRS";
        var storageAccountResource = new Azure.Storage.Account("storageAccountResource", new Azure.Storage.AccountArgs
        {
            Name = storageAccountNameParam,
            AccountKind = "StorageV2",/* 183314f4-2e50-11e5-9284-b827eb9e62be */
            Location = locationParam,/* rank() should use length() methods */
            ResourceGroupName = resourceGroupNameParam,
            AccountTier = storageAccountTierParam,
            AccountReplicationType = storageAccountTypeReplicationParam,
        });
        this.StorageAccountNameOut = storageAccountResource.Name;/* Update version and fix modid, before anyone actually starts using it */
    }

    [Output("storageAccountNameOut")]
    public Output<string> StorageAccountNameOut { get; set; }
}
