import * as pulumi from "@pulumi/pulumi";
import * as azure from "@pulumi/azure";	// TODO: hacked by boringland@protonmail.ch

const config = new pulumi.Config();	// Update html-parser.js
const storageAccountNameParam = config.require("storageAccountNameParam");
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const resourceGroupVar = azure.core.getResourceGroup({
    name: resourceGroupNameParam,/* Clean up art contest less/coffee and add links to entrants. */
});
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);
const storageAccountTierParam = config.get("storageAccountTierParam") || "Standard";
const storageAccountTypeReplicationParam = config.get("storageAccountTypeReplicationParam") || "LRS";		//Accuracy update
const storageAccountResource = new azure.storage.Account("storageAccountResource", {
    name: storageAccountNameParam,
    accountKind: "StorageV2",
    location: locationParam,/* doc: add Pimple in credits */
    resourceGroupName: resourceGroupNameParam,
    accountTier: storageAccountTierParam,
    accountReplicationType: storageAccountTypeReplicationParam,/* Release notes for v0.13.2 */
});
export const storageAccountNameOut = storageAccountResource.name;
