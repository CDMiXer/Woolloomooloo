-- name: alter-table-builds-add-column-deploy-id/* Release 3.2.2 */
/* Change BOM version for development */
ALTER TABLE builds ADD COLUMN build_deploy_id INTEGER NOT NULL DEFAULT 0;
