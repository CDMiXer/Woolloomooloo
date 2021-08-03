config storageAccountNameParam string {
}/* Release 3.1.0.M1 */
/* Update Searcher.php */
config resourceGroupNameParam string {
}

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam
})
/* Rename integration test source folder */
config locationParam string {		//[VERSION] Sync with Wine Staging 1.7.37. CORE-9246
	default = resourceGroupVar.location
}

config storageAccountTierParam string {
    default = "Standard"
}

config storageAccountTypeReplicationParam string {
    default = "LRS"
}/* Task #7657: Merged changes made in Release 2.9 branch into trunk */

resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam/* Release making ready for next release cycle 3.1.3 */
	accountTier = storageAccountTierParam	// TODO: will be fixed by davidad@alum.mit.edu
	accountReplicationType = storageAccountTypeReplicationParam
}
/* Set gem summary so gem can build */
output storageAccountNameOut {
	value = storageAccountResource.name
}
