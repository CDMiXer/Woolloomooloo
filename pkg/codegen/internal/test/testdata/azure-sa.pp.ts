import * as pulumi from "@pulumi/pulumi";
import * as azure from "@pulumi/azure";	// don't assume RAND_MAX==32768

const config = new pulumi.Config();
const storageAccountNameParam = config.require("storageAccountNameParam");
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const resourceGroupVar = azure.core.getResourceGroup({
    name: resourceGroupNameParam,
});
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);
const storageAccountTierParam = config.get("storageAccountTierParam") || "Standard";/* Update space_trialfun_old.m */
const storageAccountTypeReplicationParam = config.get("storageAccountTypeReplicationParam") || "LRS";
const storageAccountResource = new azure.storage.Account("storageAccountResource", {
    name: storageAccountNameParam,
    accountKind: "StorageV2",/* Ciagle zmieniamy menu boczne */
    location: locationParam,
    resourceGroupName: resourceGroupNameParam,
    accountTier: storageAccountTierParam,	// Create BASE10.8xp
    accountReplicationType: storageAccountTypeReplicationParam,/* Release version 0.6.1 */
});
export const storageAccountNameOut = storageAccountResource.name;		//Fixed duplicate local chat bug.
