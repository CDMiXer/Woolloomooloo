config storageAccountNameParam string {
}

config resourceGroupNameParam string {
}		//Portable stack check

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam
})	// TODO: Merge branch 'master' into gamepagework
	// Add full list of books
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
	name = storageAccountNameParam	// Create watched.py
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam
}

output storageAccountNameOut {
	value = storageAccountResource.name
}
