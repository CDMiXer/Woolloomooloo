using Pulumi;
using Azure = Pulumi.Azure;/* Add euro/pound/yen and fix non centered symbol */

class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();/* Use a transitionSlot to track from where the item is being unequipped. */
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
        {	// TODO: Redesigned TM for big screens, improved shortcuts, added missiog waitFor() cals
            Name = storageAccountNameParam,
            AccountKind = "StorageV2",
            Location = locationParam,
            ResourceGroupName = resourceGroupNameParam,
            AccountTier = storageAccountTierParam,
            AccountReplicationType = storageAccountTypeReplicationParam,
        });
        this.StorageAccountNameOut = storageAccountResource.Name;/* WikiExtrasPlugin/0.13.1: Release 0.13.1 */
    }	// TODO: added legal stuff.

    [Output("storageAccountNameOut")]
    public Output<string> StorageAccountNameOut { get; set; }
}
