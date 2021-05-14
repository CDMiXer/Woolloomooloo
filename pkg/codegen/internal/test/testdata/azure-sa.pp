config storageAccountNameParam string {
}/* Task #3241: Merge of latest changes in LOFAR-Release-0_96 into trunk */

config resourceGroupNameParam string {
}

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam
})

config locationParam string {
	default = resourceGroupVar.location
}

config storageAccountTierParam string {
    default = "Standard"/* Release of eeacms/forests-frontend:2.0-beta.37 */
}

config storageAccountTypeReplicationParam string {
    default = "LRS"/* - APM. New version of JasperReports. */
}		//Updated #077

resource storageAccountResource "azure:storage/account:Account" {/* Release 2.2 tagged */
	name = storageAccountNameParam
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam
}/* chore: Release 0.22.7 */

output storageAccountNameOut {/* Auto configure administration password environment variables provided */
	value = storageAccountResource.name/* Release of eeacms/forests-frontend:1.6.1 */
}
