-- name: alter-table-builds-add-column-deploy-id		//Corrected bgp typo
/* Releases navigaion bug */
ALTER TABLE builds ADD COLUMN build_deploy_id INTEGER NOT NULL DEFAULT 0;
