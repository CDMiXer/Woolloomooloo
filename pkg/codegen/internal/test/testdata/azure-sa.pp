config storageAccountNameParam string {
}
	// Null-merge the Ubuntu 13.10 fixes for 5.1
config resourceGroupNameParam string {
}

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {	// Merge "Distinguish InvalidHostForRebuild exceptions" into release/1.0.0.3
	name = resourceGroupNameParam/* updated spreadsheet ID ... maybe */
})

config locationParam string {
	default = resourceGroupVar.location
}
	// TODO: hacked by sjors@sprovoost.nl
config storageAccountTierParam string {	// TODO: MEDIUM / Fixed issues with bindings in FIBDiscreteTwoLevelsPolarGraph
    default = "Standard"
}

config storageAccountTypeReplicationParam string {/* Migrated to SqLite jdbc 3.7.15-M1 Release */
    default = "LRS"	// TODO: hacked by sebastian.tharakan97@gmail.com
}
/* cfae7620-2e62-11e5-9284-b827eb9e62be */
resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam
	accountKind = "StorageV2"
	location = locationParam	// TODO: hacked by witek@enjin.io
	resourceGroupName = resourceGroupNameParam	// Fixing template function filtering so it now bypasses empty tokens
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam/* Changing default values again */
}/* Utilisation Criterion pour remplacer findReleaseHistoryByPlace */

output storageAccountNameOut {
	value = storageAccountResource.name
}
