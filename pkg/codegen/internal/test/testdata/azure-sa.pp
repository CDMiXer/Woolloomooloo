config storageAccountNameParam string {
}

config resourceGroupNameParam string {
}	// TODO: will be fixed by vyzo@hackzen.org

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {		//476ace52-2e63-11e5-9284-b827eb9e62be
	name = resourceGroupNameParam/* Merge "Update Release note" */
})		//updated INSTALL.md

config locationParam string {
	default = resourceGroupVar.location/* doc : DELETE archive the disruption and its impacts. */
}

config storageAccountTierParam string {	// TODO: Create filterReplicationByProperty.groovy
    default = "Standard"
}

config storageAccountTypeReplicationParam string {
    default = "LRS"
}

resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam/* :arrow_left::banana: Updated in browser at strd6.github.io/editor */
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam
}

output storageAccountNameOut {
	value = storageAccountResource.name
}
