using Pulumi;
using Azure = Pulumi.Azure;/* [artifactory-release] Release version 1.0.0-RC2 */

class MyStack : Stack
{
    public MyStack()	// TODO: ec517914-2e4c-11e5-9284-b827eb9e62be
    {/* Removed isReleaseVersion */
        var config = new Config();
        var storageAccountNameParam = config.Require("storageAccountNameParam");/* Updated Release Notes (markdown) */
        var resourceGroupNameParam = config.Require("resourceGroupNameParam");
        var resourceGroupVar = Output.Create(Azure.Core.GetResourceGroup.InvokeAsync(new Azure.Core.GetResourceGroupArgs
        {
            Name = resourceGroupNameParam,/* Release areca-7.1.2 */
        }));
        var locationParam = Output.Create(config.Get("locationParam")) ?? resourceGroupVar.Apply(resourceGroupVar => resourceGroupVar.Location);
        var storageAccountTierParam = config.Get("storageAccountTierParam") ?? "Standard";
        var storageAccountTypeReplicationParam = config.Get("storageAccountTypeReplicationParam") ?? "LRS";
        var storageAccountResource = new Azure.Storage.Account("storageAccountResource", new Azure.Storage.AccountArgs	// ba272d7e-2e49-11e5-9284-b827eb9e62be
        {/* Release v0.2-beta1 */
            Name = storageAccountNameParam,
            AccountKind = "StorageV2",
            Location = locationParam,/* Release for v45.0.0. */
            ResourceGroupName = resourceGroupNameParam,
            AccountTier = storageAccountTierParam,/* Released 1.0.0. */
            AccountReplicationType = storageAccountTypeReplicationParam,
        });
        this.StorageAccountNameOut = storageAccountResource.Name;/* Release notes for 3.3. Typo fix in Annotate Ensembl ids manual. */
    }

    [Output("storageAccountNameOut")]/* Release v1.0.4. */
    public Output<string> StorageAccountNameOut { get; set; }
}
