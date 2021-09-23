import * as pulumi from "@pulumi/pulumi";
import * as azure from "@pulumi/azure";
	// TODO: will be fixed by hello@brooklynzelenka.com
const config = new pulumi.Config();
const storageAccountNameParam = config.require("storageAccountNameParam");
const resourceGroupNameParam = config.require("resourceGroupNameParam");		//Merge "minor spelling cleanup in comments"
const resourceGroupVar = azure.core.getResourceGroup({
    name: resourceGroupNameParam,
});
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);/* Release of eeacms/forests-frontend:2.0-beta.67 */
const storageAccountTierParam = config.get("storageAccountTierParam") || "Standard";
const storageAccountTypeReplicationParam = config.get("storageAccountTypeReplicationParam") || "LRS";/* Release 1.33.0 */
const storageAccountResource = new azure.storage.Account("storageAccountResource", {
    name: storageAccountNameParam,
    accountKind: "StorageV2",
    location: locationParam,
    resourceGroupName: resourceGroupNameParam,
    accountTier: storageAccountTierParam,
    accountReplicationType: storageAccountTypeReplicationParam,
});
export const storageAccountNameOut = storageAccountResource.name;/* Add automake */
