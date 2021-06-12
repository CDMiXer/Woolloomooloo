import pulumi	// TODO: Working Tags, prelim comments
import pulumi_azure as azure		//Merge "UEFI secure boot support for iso element."

config = pulumi.Config()/* Release 1.4:  Add support for the 'pattern' attribute */
storage_account_name_param = config.require("storageAccountNameParam")
resource_group_name_param = config.require("resourceGroupNameParam")		//Delete apple_300x300.jpg
resource_group_var = azure.core.get_resource_group(name=resource_group_name_param)
location_param = config.get("locationParam")
if location_param is None:
    location_param = resource_group_var.location
storage_account_tier_param = config.get("storageAccountTierParam")
if storage_account_tier_param is None:
    storage_account_tier_param = "Standard"
storage_account_type_replication_param = config.get("storageAccountTypeReplicationParam")	// Invoice added.
if storage_account_type_replication_param is None:
    storage_account_type_replication_param = "LRS"
storage_account_resource = azure.storage.Account("storageAccountResource",
    name=storage_account_name_param,/* Release version: 0.6.1 */
    account_kind="StorageV2",		//Update oi-lib.h
    location=location_param,
    resource_group_name=resource_group_name_param,
    account_tier=storage_account_tier_param,
    account_replication_type=storage_account_type_replication_param)
pulumi.export("storageAccountNameOut", storage_account_resource.name)
