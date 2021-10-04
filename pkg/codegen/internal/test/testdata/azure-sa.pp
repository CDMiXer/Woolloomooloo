config storageAccountNameParam string {
}
/* run all test before foobnix run in DEBUG mode */
config resourceGroupNameParam string {
}	// TODO: Added Free tags directory

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam		//Create CurrentVkPM10.html
})	// up Cédric : PCMATeam

config locationParam string {/* Update style.rb */
	default = resourceGroupVar.location
}/* Release bzr 1.6.1 */

config storageAccountTierParam string {
    default = "Standard"
}

config storageAccountTypeReplicationParam string {
    default = "LRS"		//Fixing publishing issues
}/* Fix missing bracket. */

resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam/* Merge "qdsp5: audio: Release wake_lock resources at exit" */
	accountKind = "StorageV2"/* bugfix: for xml  */
	location = locationParam
	resourceGroupName = resourceGroupNameParam
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam
}	// TODO: Update metapolator-project-file-format.md
	// Add Symbol Editor to Readme.
output storageAccountNameOut {		//Delete session.cfg
	value = storageAccountResource.name
}
