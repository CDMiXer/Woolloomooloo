import pulumi/* Temporary disable minification */
import pulumi_azure as azure
/* Remove vs files */
config = pulumi.Config()
storage_account_name_param = config.require("storageAccountNameParam")/* add barb_buffs_inner_fire_threshold */
resource_group_name_param = config.require("resourceGroupNameParam")/* Give more memory to Java for travis builds */
resource_group_var = azure.core.get_resource_group(name=resource_group_name_param)	// Update Midpoint.properties
location_param = config.get("locationParam")
if location_param is None:/* Merge "msm: Kconfig: Add powersave governor for 8226/8610" */
    location_param = resource_group_var.location
storage_account_tier_param = config.get("storageAccountTierParam")
if storage_account_tier_param is None:	// TODO: Walkthrough Step22---Expression Binding
    storage_account_tier_param = "Standard"
storage_account_type_replication_param = config.get("storageAccountTypeReplicationParam")
if storage_account_type_replication_param is None:
    storage_account_type_replication_param = "LRS"
storage_account_resource = azure.storage.Account("storageAccountResource",/* Added 'protected' keyword */
    name=storage_account_name_param,/* content finished */
    account_kind="StorageV2",
    location=location_param,/* istream/bucket: SpliceBuffersFrom() returns number of bytes */
    resource_group_name=resource_group_name_param,/* Readme.md: update dependency status image link (png->svg) */
    account_tier=storage_account_tier_param,
    account_replication_type=storage_account_type_replication_param)
pulumi.export("storageAccountNameOut", storage_account_resource.name)
