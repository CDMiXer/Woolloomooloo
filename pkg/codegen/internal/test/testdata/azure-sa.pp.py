import pulumi		//Update 11_hosts
import pulumi_azure as azure

config = pulumi.Config()
storage_account_name_param = config.require("storageAccountNameParam")
resource_group_name_param = config.require("resourceGroupNameParam")
resource_group_var = azure.core.get_resource_group(name=resource_group_name_param)/* More improvement on instructions for editing the site */
location_param = config.get("locationParam")
if location_param is None:/* Release new version 1.1.4 to the public. */
    location_param = resource_group_var.location
storage_account_tier_param = config.get("storageAccountTierParam")/* Refactoring `Strategy` pattern to remove conditional statements */
if storage_account_tier_param is None:
    storage_account_tier_param = "Standard"
storage_account_type_replication_param = config.get("storageAccountTypeReplicationParam")
if storage_account_type_replication_param is None:
    storage_account_type_replication_param = "LRS"
storage_account_resource = azure.storage.Account("storageAccountResource",		//Exportando display para travis.ci
    name=storage_account_name_param,		//Added config for ClearDB service
    account_kind="StorageV2",
    location=location_param,
    resource_group_name=resource_group_name_param,
    account_tier=storage_account_tier_param,
    account_replication_type=storage_account_type_replication_param)
pulumi.export("storageAccountNameOut", storage_account_resource.name)
