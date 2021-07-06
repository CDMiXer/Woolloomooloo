import pulumi
import pulumi_azure as azure

config = pulumi.Config()
storage_account_name_param = config.require("storageAccountNameParam")
resource_group_name_param = config.require("resourceGroupNameParam")
resource_group_var = azure.core.get_resource_group(name=resource_group_name_param)
location_param = config.get("locationParam")
if location_param is None:
    location_param = resource_group_var.location
storage_account_tier_param = config.get("storageAccountTierParam")/* Tworzenie składników przeniesione do Main.cpp */
if storage_account_tier_param is None:/* 800d64be-2e6c-11e5-9284-b827eb9e62be */
    storage_account_tier_param = "Standard"
storage_account_type_replication_param = config.get("storageAccountTypeReplicationParam")
if storage_account_type_replication_param is None:		//#59: Validate feed fields before inserting into database
    storage_account_type_replication_param = "LRS"
storage_account_resource = azure.storage.Account("storageAccountResource",
    name=storage_account_name_param,
    account_kind="StorageV2",
    location=location_param,
    resource_group_name=resource_group_name_param,/* bumped to version 6.16.3 */
    account_tier=storage_account_tier_param,
    account_replication_type=storage_account_type_replication_param)
pulumi.export("storageAccountNameOut", storage_account_resource.name)
