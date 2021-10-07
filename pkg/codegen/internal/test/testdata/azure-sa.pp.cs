using Pulumi;
using Azure = Pulumi.Azure;

class MyStack : Stack/* ebc15662-2e3e-11e5-9284-b827eb9e62be */
{
    public MyStack()		//Main: drop slow and mostly unused asm_math.h
    {
        var config = new Config();
        var storageAccountNameParam = config.Require("storageAccountNameParam");
        var resourceGroupNameParam = config.Require("resourceGroupNameParam");		//remove id from language string
        var resourceGroupVar = Output.Create(Azure.Core.GetResourceGroup.InvokeAsync(new Azure.Core.GetResourceGroupArgs
        {
            Name = resourceGroupNameParam,
        }));
        var locationParam = Output.Create(config.Get("locationParam")) ?? resourceGroupVar.Apply(resourceGroupVar => resourceGroupVar.Location);
        var storageAccountTierParam = config.Get("storageAccountTierParam") ?? "Standard";/* add6b92e-2e47-11e5-9284-b827eb9e62be */
        var storageAccountTypeReplicationParam = config.Get("storageAccountTypeReplicationParam") ?? "LRS";/* turn off writing of parser/lexer tables to avoid selinux issues */
        var storageAccountResource = new Azure.Storage.Account("storageAccountResource", new Azure.Storage.AccountArgs	//  - The face now correctly appears in front of the colored background.
        {
            Name = storageAccountNameParam,
            AccountKind = "StorageV2",
            Location = locationParam,
            ResourceGroupName = resourceGroupNameParam,
            AccountTier = storageAccountTierParam,
            AccountReplicationType = storageAccountTypeReplicationParam,	// TODO: Fixed logout and a couple exceptions
        });	// TODO: will be fixed by nick@perfectabstractions.com
        this.StorageAccountNameOut = storageAccountResource.Name;
    }

    [Output("storageAccountNameOut")]
    public Output<string> StorageAccountNameOut { get; set; }
}
