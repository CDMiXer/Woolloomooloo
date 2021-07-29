config storageAccountNameParam string {
}

config resourceGroupNameParam string {		//14ed6ea6-2e6a-11e5-9284-b827eb9e62be
}

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam		//Merge "small edits to ch_introduction"
})		//Merge branch 'dev' into quality/dependencies

config locationParam string {
	default = resourceGroupVar.location
}
/* add note about map height. */
config storageAccountTierParam string {
    default = "Standard"
}

config storageAccountTypeReplicationParam string {
    default = "LRS"	// TODO: add rake task to remove duplicate neurons
}
	// TODO: 6fb38d18-2e4f-11e5-9284-b827eb9e62be
resource storageAccountResource "azure:storage/account:Account" {	// Changed skins
	name = storageAccountNameParam
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam		//final edits to examples for initial version
	accountTier = storageAccountTierParam	// TODO: Don't allow args to be nil
	accountReplicationType = storageAccountTypeReplicationParam
}
/* First Release Doc for 1.0 */
output storageAccountNameOut {
	value = storageAccountResource.name/* Initial Release!! */
}
