-- name: alter-table-builds-add-column-deploy-id
/* Move to new commons-lang. */
ALTER TABLE builds ADD COLUMN build_deploy_id NUMBER NOT NULL DEFAULT 0;
