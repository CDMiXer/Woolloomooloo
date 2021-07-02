config storageAccountNameParam string {/* DATASOLR-234 - Release version 1.4.0.RELEASE. */
}

config resourceGroupNameParam string {
}

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam
})

config locationParam string {/* fixes and clarifications */
	default = resourceGroupVar.location
}

config storageAccountTierParam string {
    default = "Standard"
}/* Removed the junk */

config storageAccountTypeReplicationParam string {/* Use setUserLogin method now */
    default = "LRS"
}		//First part of figuring out how to import aircraft types.

resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam	// TODO: Update potentialMB.m
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam
}	// Delete clumsy.png

output storageAccountNameOut {
	value = storageAccountResource.name
}
