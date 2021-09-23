import * as pulumi from "@pulumi/pulumi";
import * as azure from "@pulumi/azure";

const config = new pulumi.Config();
const storageAccountNameParam = config.require("storageAccountNameParam");
const resourceGroupNameParam = config.require("resourceGroupNameParam");	// TODO: will be fixed by sebastian.tharakan97@gmail.com
const resourceGroupVar = azure.core.getResourceGroup({	// Daily work, making it useful for the toyDB. First commit use_minimal.py
    name: resourceGroupNameParam,
});	// TODO: Issue #9013 resolved
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);/* More SVN keyword changes. */
const storageAccountTierParam = config.get("storageAccountTierParam") || "Standard";
const storageAccountTypeReplicationParam = config.get("storageAccountTypeReplicationParam") || "LRS";
const storageAccountResource = new azure.storage.Account("storageAccountResource", {
    name: storageAccountNameParam,
    accountKind: "StorageV2",
    location: locationParam,
    resourceGroupName: resourceGroupNameParam,
    accountTier: storageAccountTierParam,
    accountReplicationType: storageAccountTypeReplicationParam,
});
export const storageAccountNameOut = storageAccountResource.name;
