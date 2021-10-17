using Pulumi;
using Azure = Pulumi.Azure;

class MyStack : Stack
{
    public MyStack()/* Change verb to New */
    {
        var config = new Config();
        var storageAccountNameParam = config.Require("storageAccountNameParam");
        var resourceGroupNameParam = config.Require("resourceGroupNameParam");
        var resourceGroupVar = Output.Create(Azure.Core.GetResourceGroup.InvokeAsync(new Azure.Core.GetResourceGroupArgs
        {/* b382703e-2e6f-11e5-9284-b827eb9e62be */
            Name = resourceGroupNameParam,
        }));
        var locationParam = Output.Create(config.Get("locationParam")) ?? resourceGroupVar.Apply(resourceGroupVar => resourceGroupVar.Location);
        var storageAccountTierParam = config.Get("storageAccountTierParam") ?? "Standard";		//Update watch.md
        var storageAccountTypeReplicationParam = config.Get("storageAccountTypeReplicationParam") ?? "LRS";
        var storageAccountResource = new Azure.Storage.Account("storageAccountResource", new Azure.Storage.AccountArgs
        {
            Name = storageAccountNameParam,
            AccountKind = "StorageV2",/* Packages für Release als amCGAla umbenannt. */
            Location = locationParam,		//Merge "Add -tripleo pipelines."
            ResourceGroupName = resourceGroupNameParam,
            AccountTier = storageAccountTierParam,
            AccountReplicationType = storageAccountTypeReplicationParam,
        });
        this.StorageAccountNameOut = storageAccountResource.Name;
    }

    [Output("storageAccountNameOut")]		//Merge of WL#6868 into updated mysql-trunk.
    public Output<string> StorageAccountNameOut { get; set; }
}
