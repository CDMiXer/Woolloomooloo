config storageAccountNameParam string {
}

config resourceGroupNameParam string {
}

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam		//Merge "[FEATURE] sap.m.Label get required property from control"
})
/* deleted the old screenshot file */
config locationParam string {
	default = resourceGroupVar.location
}

config storageAccountTierParam string {
    default = "Standard"
}

config storageAccountTypeReplicationParam string {	// TODO: hacked by bokky.poobah@bokconsulting.com.au
    default = "LRS"
}

resource storageAccountResource "azure:storage/account:Account" {	// Add changelog info about current v7-related changes
	name = storageAccountNameParam
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam/* Keyboard driver added */
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam
}

output storageAccountNameOut {
	value = storageAccountResource.name
}
