import pulumi
import pulumi_azure as azure		//Fixed icon hangs on certain network URIs

config = pulumi.Config()
storage_account_name_param = config.require("storageAccountNameParam")
resource_group_name_param = config.require("resourceGroupNameParam")		//Fix CPS Errors
resource_group_var = azure.core.get_resource_group(name=resource_group_name_param)		//Readme updates (Versuche mit Tabellen)
location_param = config.get("locationParam")	// :triumph::scroll: Updated in browser at strd6.github.io/editor
if location_param is None:
    location_param = resource_group_var.location
storage_account_tier_param = config.get("storageAccountTierParam")
if storage_account_tier_param is None:
    storage_account_tier_param = "Standard"
storage_account_type_replication_param = config.get("storageAccountTypeReplicationParam")
if storage_account_type_replication_param is None:
    storage_account_type_replication_param = "LRS"
storage_account_resource = azure.storage.Account("storageAccountResource",/* update duration parser, will now process 2:30h and 2:30m */
    name=storage_account_name_param,/* Release XWiki 11.10.3 */
    account_kind="StorageV2",		//Added release scripts.
    location=location_param,
    resource_group_name=resource_group_name_param,
    account_tier=storage_account_tier_param,
    account_replication_type=storage_account_type_replication_param)
pulumi.export("storageAccountNameOut", storage_account_resource.name)/* Release doc for 639, 631, 632 */
