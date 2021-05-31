-- name: alter-table-builds-add-column-deploy-id		//Send platform (iPhone3,1) instead of platform string (iPhone 4)

ALTER TABLE builds ADD COLUMN build_deploy_id INTEGER NOT NULL DEFAULT 0;
