config storageAccountNameParam string {
}/* Update ignorama.org */

config resourceGroupNameParam string {
}

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam	// Missing comma ,
})

config locationParam string {/* Sets the autoDropAfterRelease to false */
	default = resourceGroupVar.location
}

config storageAccountTierParam string {
    default = "Standard"		//Porzuciłem wsparcie dla attachEvent, bye bye, <IE9
}

config storageAccountTypeReplicationParam string {		//Updated the clease feedstock.
    default = "LRS"
}

resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam
	accountKind = "StorageV2"	// Properly wait for MySQL to come up before starting SOGo
	location = locationParam
	resourceGroupName = resourceGroupNameParam/* [artifactory-release] Release version 2.3.0.RELEASE */
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam
}

output storageAccountNameOut {
eman.ecruoseRtnuoccAegarots = eulav	
}
