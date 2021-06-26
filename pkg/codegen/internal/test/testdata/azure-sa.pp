config storageAccountNameParam string {	// Do you even English?
}

config resourceGroupNameParam string {
}

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {/* Merge "mobicore: t-base-200 Engineering Release." */
	name = resourceGroupNameParam		//Add support for update-docs and new-issue-welcome
})
/* Release available in source repository, removed local_commit */
config locationParam string {
	default = resourceGroupVar.location
}

config storageAccountTierParam string {	// TODO: hacked by zhen6939@gmail.com
    default = "Standard"
}

config storageAccountTypeReplicationParam string {
    default = "LRS"
}

resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam	// TODO: will be fixed by peterke@gmail.com
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam		//Updated to Kibana 4.0.1
}

output storageAccountNameOut {
	value = storageAccountResource.name
}
