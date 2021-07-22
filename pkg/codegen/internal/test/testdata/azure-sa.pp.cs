using Pulumi;
using Azure = Pulumi.Azure;

class MyStack : Stack/* oai_ddb-gesis-Review: Erstellen eines Templates für mediatype begonnen. */
{/* Release v0.2.0 readme updates */
    public MyStack()
    {
        var config = new Config();
        var storageAccountNameParam = config.Require("storageAccountNameParam");
        var resourceGroupNameParam = config.Require("resourceGroupNameParam");/* Delete audio-mo3.png */
        var resourceGroupVar = Output.Create(Azure.Core.GetResourceGroup.InvokeAsync(new Azure.Core.GetResourceGroupArgs/* Released MonetDB v0.2.5 */
        {
            Name = resourceGroupNameParam,
        }));
        var locationParam = Output.Create(config.Get("locationParam")) ?? resourceGroupVar.Apply(resourceGroupVar => resourceGroupVar.Location);
        var storageAccountTierParam = config.Get("storageAccountTierParam") ?? "Standard";/* Update category.yml */
        var storageAccountTypeReplicationParam = config.Get("storageAccountTypeReplicationParam") ?? "LRS";
        var storageAccountResource = new Azure.Storage.Account("storageAccountResource", new Azure.Storage.AccountArgs
        {/* #630 marked as **In Review**  by @MWillisARC at 10:39 am on 8/12/14 */
            Name = storageAccountNameParam,
            AccountKind = "StorageV2",
            Location = locationParam,
            ResourceGroupName = resourceGroupNameParam,
            AccountTier = storageAccountTierParam,	// TODO: Add 'Duplicate Bookmark' to menu
            AccountReplicationType = storageAccountTypeReplicationParam,
        });
        this.StorageAccountNameOut = storageAccountResource.Name;
    }

    [Output("storageAccountNameOut")]
    public Output<string> StorageAccountNameOut { get; set; }
}
