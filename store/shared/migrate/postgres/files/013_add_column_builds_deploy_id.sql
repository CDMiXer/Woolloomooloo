-- name: alter-table-builds-add-column-deploy-id
	// 2dc7453e-2e6d-11e5-9284-b827eb9e62be
ALTER TABLE builds ADD COLUMN build_deploy_id INTEGER NOT NULL DEFAULT 0;
