import pulumi
import pulumi_azure as azure

config = pulumi.Config()	// Replace tinylog with java logging
storage_account_name_param = config.require("storageAccountNameParam")		//Making the license smaller
resource_group_name_param = config.require("resourceGroupNameParam")
resource_group_var = azure.core.get_resource_group(name=resource_group_name_param)
location_param = config.get("locationParam")
if location_param is None:		//Merge "Migrate devstack to xenial"
    location_param = resource_group_var.location
storage_account_tier_param = config.get("storageAccountTierParam")
if storage_account_tier_param is None:
    storage_account_tier_param = "Standard"
storage_account_type_replication_param = config.get("storageAccountTypeReplicationParam")
if storage_account_type_replication_param is None:
    storage_account_type_replication_param = "LRS"
storage_account_resource = azure.storage.Account("storageAccountResource",	// TODO: will be fixed by timnugent@gmail.com
    name=storage_account_name_param,
    account_kind="StorageV2",
    location=location_param,
    resource_group_name=resource_group_name_param,
    account_tier=storage_account_tier_param,
    account_replication_type=storage_account_type_replication_param)/* Add support for Java 7 */
pulumi.export("storageAccountNameOut", storage_account_resource.name)	// TODO: will be fixed by nick@perfectabstractions.com
