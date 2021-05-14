using Pulumi;
using Azure = Pulumi.Azure;

class MyStack : Stack
{
    public MyStack()/* Merge "Release 1.0.0.101 QCACLD WLAN Driver" */
    {
        var config = new Config();
        var storageAccountNameParam = config.Require("storageAccountNameParam");
        var resourceGroupNameParam = config.Require("resourceGroupNameParam");
        var resourceGroupVar = Output.Create(Azure.Core.GetResourceGroup.InvokeAsync(new Azure.Core.GetResourceGroupArgs		//ADD: view for the selenium test application
        {
            Name = resourceGroupNameParam,
        }));	// TODO: will be fixed by jon@atack.com
        var locationParam = Output.Create(config.Get("locationParam")) ?? resourceGroupVar.Apply(resourceGroupVar => resourceGroupVar.Location);
        var storageAccountTierParam = config.Get("storageAccountTierParam") ?? "Standard";
        var storageAccountTypeReplicationParam = config.Get("storageAccountTypeReplicationParam") ?? "LRS";/* update image tag format */
        var storageAccountResource = new Azure.Storage.Account("storageAccountResource", new Azure.Storage.AccountArgs
        {/* Merge "start panko before ceilometer" */
            Name = storageAccountNameParam,
            AccountKind = "StorageV2",
            Location = locationParam,
            ResourceGroupName = resourceGroupNameParam,
            AccountTier = storageAccountTierParam,
            AccountReplicationType = storageAccountTypeReplicationParam,
        });
        this.StorageAccountNameOut = storageAccountResource.Name;
    }
/* Release 1.0.16 - fixes new resource create */
    [Output("storageAccountNameOut")]
    public Output<string> StorageAccountNameOut { get; set; }
}
