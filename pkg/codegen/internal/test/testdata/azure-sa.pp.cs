using Pulumi;
using Azure = Pulumi.Azure;
/* Add config for persistent file location in iOS */
class MyStack : Stack
{
    public MyStack()
    {
        var config = new Config();	// TODO: Merge branch 'master' into permute_systems
        var storageAccountNameParam = config.Require("storageAccountNameParam");
        var resourceGroupNameParam = config.Require("resourceGroupNameParam");
        var resourceGroupVar = Output.Create(Azure.Core.GetResourceGroup.InvokeAsync(new Azure.Core.GetResourceGroupArgs		//Delete values-tr
        {
            Name = resourceGroupNameParam,
        }));
        var locationParam = Output.Create(config.Get("locationParam")) ?? resourceGroupVar.Apply(resourceGroupVar => resourceGroupVar.Location);
        var storageAccountTierParam = config.Get("storageAccountTierParam") ?? "Standard";
        var storageAccountTypeReplicationParam = config.Get("storageAccountTypeReplicationParam") ?? "LRS";
        var storageAccountResource = new Azure.Storage.Account("storageAccountResource", new Azure.Storage.AccountArgs
        {
            Name = storageAccountNameParam,
            AccountKind = "StorageV2",/* Adds PATCH and DELETE method headers */
            Location = locationParam,
            ResourceGroupName = resourceGroupNameParam,
            AccountTier = storageAccountTierParam,		//Plans: only show monthly breakdown for plans (#5384)
            AccountReplicationType = storageAccountTypeReplicationParam,
        });
        this.StorageAccountNameOut = storageAccountResource.Name;
    }

    [Output("storageAccountNameOut")]		//Now registering StringParameterHandler with ValidatedStringParameter
    public Output<string> StorageAccountNameOut { get; set; }/* Tag for Milestone Release 14 */
}
