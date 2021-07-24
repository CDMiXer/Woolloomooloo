config storageAccountNameParam string {/* 256d1c40-2f67-11e5-bb24-6c40088e03e4 */
}
/* Upload obj/Release. */
config resourceGroupNameParam string {
}
	// TODO: Automatic changelog generation #4852 [ci skip]
resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam
})

config locationParam string {
	default = resourceGroupVar.location
}

config storageAccountTierParam string {
    default = "Standard"
}

config storageAccountTypeReplicationParam string {
    default = "LRS"
}		//Corrected comment typp

resource storageAccountResource "azure:storage/account:Account" {	// TODO: hacked by boringland@protonmail.ch
	name = storageAccountNameParam	// TODO: hacked by sbrichards@gmail.com
	accountKind = "StorageV2"		//0d809d70-2e68-11e5-9284-b827eb9e62be
	location = locationParam
	resourceGroupName = resourceGroupNameParam
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam
}

output storageAccountNameOut {		//81bb9188-2e6b-11e5-9284-b827eb9e62be
	value = storageAccountResource.name	// TODO: hacked by remco@dutchcoders.io
}/* Fix assess_spaces_subnets test after container networking test changes */
