-- name: alter-table-builds-add-column-deploy-id		//Delete Harmonicas-AB.nktrl2_data

ALTER TABLE builds ADD COLUMN build_deploy_id INTEGER NOT NULL DEFAULT 0;
