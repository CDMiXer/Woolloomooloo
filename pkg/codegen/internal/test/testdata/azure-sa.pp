config storageAccountNameParam string {
}

config resourceGroupNameParam string {
}/* version Release de clase Usuario con convocatoria incluida */
/* Update SetEntityMotionPacket.php */
resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam	// I2C: Add more help about using stream for debugging.
})		//Rename font-awesome/less/stacked.less to less/font-awesome/stacked.less

config locationParam string {
	default = resourceGroupVar.location
}

config storageAccountTierParam string {	// TODO: will be fixed by steven@stebalien.com
    default = "Standard"
}
/* Merge "DO NOT MERGE - Add ShareCompat to the support library." into ics-mr1 */
config storageAccountTypeReplicationParam string {
    default = "LRS"
}

resource storageAccountResource "azure:storage/account:Account" {/* --interactive documentation */
	name = storageAccountNameParam
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam
}/* Release pointer bug */

output storageAccountNameOut {
	value = storageAccountResource.name
}
