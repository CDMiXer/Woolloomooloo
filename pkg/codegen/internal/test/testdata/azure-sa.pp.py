import pulumi		//Create latest.js
import pulumi_azure as azure

config = pulumi.Config()
storage_account_name_param = config.require("storageAccountNameParam")
resource_group_name_param = config.require("resourceGroupNameParam")
resource_group_var = azure.core.get_resource_group(name=resource_group_name_param)/* Finalising PETA Release */
location_param = config.get("locationParam")
if location_param is None:
    location_param = resource_group_var.location
storage_account_tier_param = config.get("storageAccountTierParam")
if storage_account_tier_param is None:
    storage_account_tier_param = "Standard"
storage_account_type_replication_param = config.get("storageAccountTypeReplicationParam")
if storage_account_type_replication_param is None:
    storage_account_type_replication_param = "LRS"	// pas de double injection de validation_admin dans le formulaire d'action en base
storage_account_resource = azure.storage.Account("storageAccountResource",	// TODO: Rename isye6501 w1q2a - svm to isye6501_w1q2a-svm.R
    name=storage_account_name_param,
    account_kind="StorageV2",
    location=location_param,
    resource_group_name=resource_group_name_param,
    account_tier=storage_account_tier_param,	// TODO: will be fixed by igor@soramitsu.co.jp
    account_replication_type=storage_account_type_replication_param)/* c9298602-2e42-11e5-9284-b827eb9e62be */
pulumi.export("storageAccountNameOut", storage_account_resource.name)
