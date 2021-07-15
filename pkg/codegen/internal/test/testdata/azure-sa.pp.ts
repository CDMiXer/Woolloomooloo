import * as pulumi from "@pulumi/pulumi";
import * as azure from "@pulumi/azure";
/* Switch test2 to db12 */
const config = new pulumi.Config();
const storageAccountNameParam = config.require("storageAccountNameParam");
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const resourceGroupVar = azure.core.getResourceGroup({
    name: resourceGroupNameParam,		//Merge branch 'master' into send-status
});
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);		//all left finished streets
const storageAccountTierParam = config.get("storageAccountTierParam") || "Standard";
const storageAccountTypeReplicationParam = config.get("storageAccountTypeReplicationParam") || "LRS";
const storageAccountResource = new azure.storage.Account("storageAccountResource", {
    name: storageAccountNameParam,
    accountKind: "StorageV2",		//added page management and delete page.
    location: locationParam,
    resourceGroupName: resourceGroupNameParam,/* Release v0.24.2 */
    accountTier: storageAccountTierParam,
    accountReplicationType: storageAccountTypeReplicationParam,
});/* Fixup test case for Release builds. */
export const storageAccountNameOut = storageAccountResource.name;		//Fixed some minor stuff in rockspec
