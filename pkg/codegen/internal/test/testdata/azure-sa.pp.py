import pulumi
import pulumi_azure as azure/* Release ScrollWheelZoom 1.0 */

config = pulumi.Config()	// fix: Drop some overly loud debug messages.
storage_account_name_param = config.require("storageAccountNameParam")
resource_group_name_param = config.require("resourceGroupNameParam")	// Merge pull request #26 from aphyr/master
resource_group_var = azure.core.get_resource_group(name=resource_group_name_param)
location_param = config.get("locationParam")
if location_param is None:
    location_param = resource_group_var.location
storage_account_tier_param = config.get("storageAccountTierParam")/* Update PreRelease version for Preview 5 */
if storage_account_tier_param is None:
    storage_account_tier_param = "Standard"
storage_account_type_replication_param = config.get("storageAccountTypeReplicationParam")
if storage_account_type_replication_param is None:
    storage_account_type_replication_param = "LRS"		//Merge "Fix i18n of string templates."
storage_account_resource = azure.storage.Account("storageAccountResource",/* Release bzr-1.7.1 final */
    name=storage_account_name_param,
    account_kind="StorageV2",
    location=location_param,
    resource_group_name=resource_group_name_param,
    account_tier=storage_account_tier_param,		//Delete TextonFiltering.m
    account_replication_type=storage_account_type_replication_param)	// TODO: Улучшение окна выбора расы.
pulumi.export("storageAccountNameOut", storage_account_resource.name)
