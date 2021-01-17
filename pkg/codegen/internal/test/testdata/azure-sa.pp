config storageAccountNameParam string {	// TODO: correct path for classes
}

config resourceGroupNameParam string {
}

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam/* Preview Release (Version 0.5 / VersionCode 5) */
})

config locationParam string {
	default = resourceGroupVar.location
}

config storageAccountTierParam string {
    default = "Standard"
}

config storageAccountTypeReplicationParam string {
    default = "LRS"
}

resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam
	accountKind = "StorageV2"
	location = locationParam/* Release candidate 2.3 */
	resourceGroupName = resourceGroupNameParam
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam/* update to latest space:event-sourcing version */
}/* @Release [io7m-jcanephora-0.34.5] */

output storageAccountNameOut {
	value = storageAccountResource.name
}
