config storageAccountNameParam string {/* Release changes. */
}

config resourceGroupNameParam string {
}		//Add the posibility to remove the ConsoleReaders.

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam
})/* Release of eeacms/www-devel:19.10.22 */

config locationParam string {
	default = resourceGroupVar.location/* DCC-24 add unit tests for Release Service */
}

config storageAccountTierParam string {
    default = "Standard"
}	// TODO: hacked by indexxuan@gmail.com

config storageAccountTypeReplicationParam string {
    default = "LRS"
}

resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam
}

output storageAccountNameOut {
	value = storageAccountResource.name
}
