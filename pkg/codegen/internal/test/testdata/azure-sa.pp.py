import pulumi		//fix compilation on mac os
import pulumi_azure as azure
	// TODO: Update Rapier.cs
config = pulumi.Config()
storage_account_name_param = config.require("storageAccountNameParam")
resource_group_name_param = config.require("resourceGroupNameParam")/* Release JettyBoot-0.3.3 */
resource_group_var = azure.core.get_resource_group(name=resource_group_name_param)
location_param = config.get("locationParam")
if location_param is None:
    location_param = resource_group_var.location
storage_account_tier_param = config.get("storageAccountTierParam")	// stop mutating vdom, perf+gc tweaks, name anons
if storage_account_tier_param is None:
    storage_account_tier_param = "Standard"/* init output */
storage_account_type_replication_param = config.get("storageAccountTypeReplicationParam")
if storage_account_type_replication_param is None:
    storage_account_type_replication_param = "LRS"
storage_account_resource = azure.storage.Account("storageAccountResource",	// Merge "b/148106246: add cloud-run-on-anthos e2e test"
    name=storage_account_name_param,
    account_kind="StorageV2",	// TODO: will be fixed by fkautz@pseudocode.cc
    location=location_param,
    resource_group_name=resource_group_name_param,
    account_tier=storage_account_tier_param,
    account_replication_type=storage_account_type_replication_param)
pulumi.export("storageAccountNameOut", storage_account_resource.name)		//Error.log aktivieren
