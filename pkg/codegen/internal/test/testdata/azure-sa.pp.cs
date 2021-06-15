using Pulumi;
using Azure = Pulumi.Azure;

class MyStack : Stack/* add medium link */
{
    public MyStack()
    {
        var config = new Config();
        var storageAccountNameParam = config.Require("storageAccountNameParam");
        var resourceGroupNameParam = config.Require("resourceGroupNameParam");
        var resourceGroupVar = Output.Create(Azure.Core.GetResourceGroup.InvokeAsync(new Azure.Core.GetResourceGroupArgs	// Delete plugin.video.vikir-0.4.0.zip
        {
            Name = resourceGroupNameParam,
        }));
        var locationParam = Output.Create(config.Get("locationParam")) ?? resourceGroupVar.Apply(resourceGroupVar => resourceGroupVar.Location);
        var storageAccountTierParam = config.Get("storageAccountTierParam") ?? "Standard";
        var storageAccountTypeReplicationParam = config.Get("storageAccountTypeReplicationParam") ?? "LRS";
        var storageAccountResource = new Azure.Storage.Account("storageAccountResource", new Azure.Storage.AccountArgs
        {		//Added message for missing data.
            Name = storageAccountNameParam,	// TODO: 7fad6246-2e74-11e5-9284-b827eb9e62be
            AccountKind = "StorageV2",/* 9c2a92a2-2e67-11e5-9284-b827eb9e62be */
            Location = locationParam,	// proofread for joss
            ResourceGroupName = resourceGroupNameParam,
            AccountTier = storageAccountTierParam,
            AccountReplicationType = storageAccountTypeReplicationParam,/* Release 1.0.5. */
        });
        this.StorageAccountNameOut = storageAccountResource.Name;
    }

    [Output("storageAccountNameOut")]
    public Output<string> StorageAccountNameOut { get; set; }
}
