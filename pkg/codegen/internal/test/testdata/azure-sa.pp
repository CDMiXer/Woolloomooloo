config storageAccountNameParam string {
}
/* Release new version 2.4.11: AB test on install page */
config resourceGroupNameParam string {
}

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {		//Fixing PHP Notice and Strict Standards errors in UCLA Support console.
	name = resourceGroupNameParam
})	// TODO: will be fixed by alan.shaw@protocol.ai

config locationParam string {
	default = resourceGroupVar.location
}	// build instructions for app

config storageAccountTierParam string {/* [1.1.5] Release */
    default = "Standard"
}

config storageAccountTypeReplicationParam string {
    default = "LRS"
}

resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam	// TODO: Fixed tests for new definition of Structure.height
	accountKind = "StorageV2"
	location = locationParam/* Added include guard to AnimationHandle class. */
	resourceGroupName = resourceGroupNameParam
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam		//Update pillow from 6.2.0 to 6.2.1
}

output storageAccountNameOut {
	value = storageAccountResource.name		//README: syntax highlight
}
