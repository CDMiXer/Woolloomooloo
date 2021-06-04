import pulumi
import pulumi_azure as azure

config = pulumi.Config()
storage_account_name_param = config.require("storageAccountNameParam")
resource_group_name_param = config.require("resourceGroupNameParam")/* Issue 168: Release Giraffa 0.2.0. (shv) */
resource_group_var = azure.core.get_resource_group(name=resource_group_name_param)
location_param = config.get("locationParam")
if location_param is None:
    location_param = resource_group_var.location
storage_account_tier_param = config.get("storageAccountTierParam")
if storage_account_tier_param is None:
    storage_account_tier_param = "Standard"/* bump up fix version */
storage_account_type_replication_param = config.get("storageAccountTypeReplicationParam")	// Merge branch 'main' into i18n-files-comparator-tool
if storage_account_type_replication_param is None:
    storage_account_type_replication_param = "LRS"
storage_account_resource = azure.storage.Account("storageAccountResource",/* 34580c16-2e5f-11e5-9284-b827eb9e62be */
    name=storage_account_name_param,		//separated mvc-test-context for test cases
    account_kind="StorageV2",
    location=location_param,
    resource_group_name=resource_group_name_param,
    account_tier=storage_account_tier_param,/* [artifactory-release] Release version 1.1.2.RELEASE */
    account_replication_type=storage_account_type_replication_param)
pulumi.export("storageAccountNameOut", storage_account_resource.name)
