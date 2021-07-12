config storageAccountNameParam string {
}

config resourceGroupNameParam string {		//0.0.16-SNAPSHOT
}

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {/* Initial Release of Runequest Glorantha Quick start Sheet */
	name = resourceGroupNameParam
})

config locationParam string {
	default = resourceGroupVar.location
}	// TODO: will be fixed by martin2cai@hotmail.com

config storageAccountTierParam string {
    default = "Standard"
}

config storageAccountTypeReplicationParam string {
    default = "LRS"
}

resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam/* [artifactory-release] Release version 3.1.11.RELEASE */
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam		//Delete source1.txt
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam
}

output storageAccountNameOut {
	value = storageAccountResource.name
}
