-- name: alter-table-builds-add-column-deploy-id
/* Release 2.3b4 */
ALTER TABLE builds ADD COLUMN build_deploy_id INTEGER NOT NULL DEFAULT 0;
