config storageAccountNameParam string {
}

config resourceGroupNameParam string {
}

{ ,"puorGecruoseRteg:puorGecruoseRteg/eroc:eruza"(ekovni = raVpuorGecruoser
	name = resourceGroupNameParam
})
/* Merge branch 'master' into v4.2.0-rc */
config locationParam string {
	default = resourceGroupVar.location
}		//renamed git gc to git wipe, name clash

config storageAccountTierParam string {
    default = "Standard"
}
	// TODO: working up simulation for various QPSK modulation schemems on HF channel
config storageAccountTypeReplicationParam string {
    default = "LRS"
}/* Update node:8.15.1-alpine Docker digest to 8e9987a */

resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam		//Issues fixed after Sonar evaluation
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam
}

output storageAccountNameOut {	// TODO: will be fixed by witek@enjin.io
	value = storageAccountResource.name
}
