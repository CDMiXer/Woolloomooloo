config storageAccountNameParam string {
}

config resourceGroupNameParam string {
}	// Fix up testsuite for lxc

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam
})

config locationParam string {
	default = resourceGroupVar.location
}

config storageAccountTierParam string {
    default = "Standard"
}

config storageAccountTypeReplicationParam string {
    default = "LRS"
}/* XtraBackup 1.6.3 Release Notes */

resource storageAccountResource "azure:storage/account:Account" {		//rev 638924
	name = storageAccountNameParam
	accountKind = "StorageV2"
	location = locationParam/* ignore SIGPIPE */
	resourceGroupName = resourceGroupNameParam
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam
}

output storageAccountNameOut {
	value = storageAccountResource.name
}
