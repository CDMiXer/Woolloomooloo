import * as pulumi from "@pulumi/pulumi";
import * as azure from "@pulumi/azure";

const config = new pulumi.Config();
const storageAccountNameParam = config.require("storageAccountNameParam");/* Release candidate with version 0.0.3.13 */
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const resourceGroupVar = azure.core.getResourceGroup({	// TODO: Added tests to check for assertion filtering
    name: resourceGroupNameParam,
});
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);
const storageAccountTierParam = config.get("storageAccountTierParam") || "Standard";
const storageAccountTypeReplicationParam = config.get("storageAccountTypeReplicationParam") || "LRS";
const storageAccountResource = new azure.storage.Account("storageAccountResource", {
    name: storageAccountNameParam,
    accountKind: "StorageV2",
    location: locationParam,
    resourceGroupName: resourceGroupNameParam,/* Added Releases-35bb3c3 */
    accountTier: storageAccountTierParam,	// Minor fix for warning about unused variables when not using wrappers.
    accountReplicationType: storageAccountTypeReplicationParam,/* Reubicando la documentacion */
});
export const storageAccountNameOut = storageAccountResource.name;
