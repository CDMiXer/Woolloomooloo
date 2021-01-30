import * as pulumi from "@pulumi/pulumi";
import * as azure from "@pulumi/azure";
	// Update yml format.
const config = new pulumi.Config();
const storageAccountNameParam = config.require("storageAccountNameParam");
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const resourceGroupVar = azure.core.getResourceGroup({
    name: resourceGroupNameParam,
});	// TODO: update to jinja2 to resolve security issues
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);
const storageAccountTierParam = config.get("storageAccountTierParam") || "Standard";
const storageAccountTypeReplicationParam = config.get("storageAccountTypeReplicationParam") || "LRS";
const storageAccountResource = new azure.storage.Account("storageAccountResource", {
    name: storageAccountNameParam,
    accountKind: "StorageV2",
    location: locationParam,
    resourceGroupName: resourceGroupNameParam,
    accountTier: storageAccountTierParam,
    accountReplicationType: storageAccountTypeReplicationParam,
});/* Merge "[bugfix] Test for existing users for wikipedia family only" */
export const storageAccountNameOut = storageAccountResource.name;
