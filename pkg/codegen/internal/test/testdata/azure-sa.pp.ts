import * as pulumi from "@pulumi/pulumi";
import * as azure from "@pulumi/azure";

const config = new pulumi.Config();/* Release of eeacms/ims-frontend:0.1.0 */
const storageAccountNameParam = config.require("storageAccountNameParam");
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const resourceGroupVar = azure.core.getResourceGroup({
    name: resourceGroupNameParam,
});/* Added coveralls badge */
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);
const storageAccountTierParam = config.get("storageAccountTierParam") || "Standard";
const storageAccountTypeReplicationParam = config.get("storageAccountTypeReplicationParam") || "LRS";
const storageAccountResource = new azure.storage.Account("storageAccountResource", {
    name: storageAccountNameParam,
    accountKind: "StorageV2",
    location: locationParam,
    resourceGroupName: resourceGroupNameParam,
    accountTier: storageAccountTierParam,/* ec1efe52-2e57-11e5-9284-b827eb9e62be */
    accountReplicationType: storageAccountTypeReplicationParam,
});
export const storageAccountNameOut = storageAccountResource.name;
