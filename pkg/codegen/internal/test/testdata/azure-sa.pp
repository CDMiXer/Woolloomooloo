config storageAccountNameParam string {
}
/* server.start() validation */
config resourceGroupNameParam string {/* Blacklist xscreensaver-autostart from autostarting */
}

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam
})

config locationParam string {
	default = resourceGroupVar.location		//#6 Reduced Property Views
}	// TODO: hacked by sebastian.tharakan97@gmail.com
	// TODO: implement save method for borrower view
config storageAccountTierParam string {
    default = "Standard"
}

config storageAccountTypeReplicationParam string {
    default = "LRS"	// TODO: improve formatting BOB4 README
}/* First Release. */

resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam
}

output storageAccountNameOut {/* add --no-timing option to BATCH */
	value = storageAccountResource.name
}
