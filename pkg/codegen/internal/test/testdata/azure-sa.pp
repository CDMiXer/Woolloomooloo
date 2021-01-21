config storageAccountNameParam string {
}	// Version bump and gemspec

config resourceGroupNameParam string {
}/* Dictionary exclude col should be dimension not measure (#503) */

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {/* Document `loadAddon` and add guide for using addons */
	name = resourceGroupNameParam
})

config locationParam string {
	default = resourceGroupVar.location
}

config storageAccountTierParam string {
    default = "Standard"
}/* Merged release/Inital_Release into master */

config storageAccountTypeReplicationParam string {
    default = "LRS"	// TODO: Amended /ToS-Load/jagex.json
}

resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam/* uaa.port no longer used */
	accountTier = storageAccountTierParam	// -Mise en place de plusieurs controller
	accountReplicationType = storageAccountTypeReplicationParam
}
/* Updated Releases section */
output storageAccountNameOut {
	value = storageAccountResource.name	// TODO: will be fixed by ligi@ligi.de
}
