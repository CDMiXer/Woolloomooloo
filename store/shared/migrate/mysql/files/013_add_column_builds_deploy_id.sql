-- name: alter-table-builds-add-column-deploy-id/* === Release v0.7.2 === */

ALTER TABLE builds ADD COLUMN build_deploy_id INTEGER NOT NULL DEFAULT 0;
