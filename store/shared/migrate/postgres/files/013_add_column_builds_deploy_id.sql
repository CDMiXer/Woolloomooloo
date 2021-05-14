-- name: alter-table-builds-add-column-deploy-id
		//Add tests to cover code exposed by removal of external compilation
ALTER TABLE builds ADD COLUMN build_deploy_id INTEGER NOT NULL DEFAULT 0;
