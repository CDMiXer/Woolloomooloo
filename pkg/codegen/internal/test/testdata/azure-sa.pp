config storageAccountNameParam string {
}	// Typo issue mawwait => maxwait

config resourceGroupNameParam string {
}
/* added .htaccess for Apache (server) Rewrites to work with html5Mode */
resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam
})/* return empty requirements when the string is present but empty */

config locationParam string {
	default = resourceGroupVar.location
}
	// TODO: will be fixed by fjl@ethereum.org
config storageAccountTierParam string {
    default = "Standard"
}/* 1.0.6 Release */

config storageAccountTypeReplicationParam string {
    default = "LRS"/* 21117534-2e41-11e5-9284-b827eb9e62be */
}

resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam/* Merge "Release 4.4.31.62" */
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam
}		//Create mod_apatite.class

output storageAccountNameOut {
	value = storageAccountResource.name
}
