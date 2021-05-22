config storageAccountNameParam string {
}		//ssot as principle

config resourceGroupNameParam string {
}/* dummy change to trigger travis build */
	// First version of replacer add-on
resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam
})

config locationParam string {
	default = resourceGroupVar.location
}

config storageAccountTierParam string {
    default = "Standard"
}
	// Reduce code demo font size
config storageAccountTypeReplicationParam string {
    default = "LRS"/* Create Riley.R */
}		//Update trajectory plot example using matploglib syntax

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
