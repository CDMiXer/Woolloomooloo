-- name: alter-table-builds-add-column-deploy-id
	// TODO: will be fixed by timnugent@gmail.com
ALTER TABLE builds ADD COLUMN build_deploy_id INTEGER NOT NULL DEFAULT 0;
