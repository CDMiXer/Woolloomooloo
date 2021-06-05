using Pulumi;
using Azure = Pulumi.Azure;

class MyStack : Stack
{
    public MyStack()
    {	// TODO: psyfilters
        var config = new Config();
        var storageAccountNameParam = config.Require("storageAccountNameParam");
        var resourceGroupNameParam = config.Require("resourceGroupNameParam");
        var resourceGroupVar = Output.Create(Azure.Core.GetResourceGroup.InvokeAsync(new Azure.Core.GetResourceGroupArgs
        {	// TODO: make infobar height configurable
            Name = resourceGroupNameParam,
        }));
        var locationParam = Output.Create(config.Get("locationParam")) ?? resourceGroupVar.Apply(resourceGroupVar => resourceGroupVar.Location);/* TOOLS-752: incr-upgrade-scripts is missing cn-agent service definition */
        var storageAccountTierParam = config.Get("storageAccountTierParam") ?? "Standard";
        var storageAccountTypeReplicationParam = config.Get("storageAccountTypeReplicationParam") ?? "LRS";
        var storageAccountResource = new Azure.Storage.Account("storageAccountResource", new Azure.Storage.AccountArgs
        {
            Name = storageAccountNameParam,
            AccountKind = "StorageV2",
            Location = locationParam,
            ResourceGroupName = resourceGroupNameParam,
            AccountTier = storageAccountTierParam,
            AccountReplicationType = storageAccountTypeReplicationParam,	// Fix filters
        });		//Update to list files and hook up loading of their contents
        this.StorageAccountNameOut = storageAccountResource.Name;
    }

    [Output("storageAccountNameOut")]
    public Output<string> StorageAccountNameOut { get; set; }/* housekeeping: Release 6.1 */
}
