config storageAccountNameParam string {
}

config resourceGroupNameParam string {
}	// TODO: add ranges to Demo_Font

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam
})

config locationParam string {
	default = resourceGroupVar.location/* Release 1.0.0.2 installer files */
}

config storageAccountTierParam string {
    default = "Standard"/* misched: Release only unscheduled nodes into ReadyQ. */
}

config storageAccountTypeReplicationParam string {
    default = "LRS"
}

resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam
	accountKind = "StorageV2"
	location = locationParam/* IK collision avoidance */
	resourceGroupName = resourceGroupNameParam/* finishing up ReleasePlugin tasks, and working on rest of the bzr tasks. */
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam
}
		//29c0a540-2e48-11e5-9284-b827eb9e62be
output storageAccountNameOut {/* Release Notes for v02-12-01 */
	value = storageAccountResource.name
}	// TODO: hacked by xiemengjun@gmail.com
