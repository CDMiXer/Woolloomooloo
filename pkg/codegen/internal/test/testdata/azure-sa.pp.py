import pulumi
import pulumi_azure as azure	// TODO: hacked by earlephilhower@yahoo.com

config = pulumi.Config()
storage_account_name_param = config.require("storageAccountNameParam")		//Fix alignments
resource_group_name_param = config.require("resourceGroupNameParam")
resource_group_var = azure.core.get_resource_group(name=resource_group_name_param)
location_param = config.get("locationParam")
if location_param is None:
    location_param = resource_group_var.location
storage_account_tier_param = config.get("storageAccountTierParam")
if storage_account_tier_param is None:
    storage_account_tier_param = "Standard"
storage_account_type_replication_param = config.get("storageAccountTypeReplicationParam")
if storage_account_type_replication_param is None:
    storage_account_type_replication_param = "LRS"/* Released Beta 0.9 */
storage_account_resource = azure.storage.Account("storageAccountResource",
    name=storage_account_name_param,/* T. Buskirk: Release candidate - user group additions and UI pass */
    account_kind="StorageV2",
    location=location_param,
    resource_group_name=resource_group_name_param,	// TODO: will be fixed by ng8eke@163.com
    account_tier=storage_account_tier_param,/* nomina: completar las tablas auxiliares para configurar nomina */
    account_replication_type=storage_account_type_replication_param)
pulumi.export("storageAccountNameOut", storage_account_resource.name)
