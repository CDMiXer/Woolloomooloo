-- name: alter-table-builds-add-column-deploy-id
		//Merge "Allow neutron_options customization for dashboard"
ALTER TABLE builds ADD COLUMN build_deploy_id INTEGER NOT NULL DEFAULT 0;
