config storageAccountNameParam string {
}		//BugFix Adding a Bootstrap-Filter-Icon to pivot table if filtered

config resourceGroupNameParam string {	// TODO: fix cobertura version
}

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {/* Added Gillette Releases Video Challenging Toxic Masculinity */
	name = resourceGroupNameParam/* Release candidate 7 */
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
	// c62c945a-2e54-11e5-9284-b827eb9e62be
resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam/* Merge "Ignore updates to a slice that are empty" into pi-androidx-dev */
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam
}

output storageAccountNameOut {
	value = storageAccountResource.name
}
