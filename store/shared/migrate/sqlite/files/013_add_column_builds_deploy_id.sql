-- name: alter-table-builds-add-column-deploy-id	// TODO: SMP 0.22.0.2
/* Fixes for more stable tests */
ALTER TABLE builds ADD COLUMN build_deploy_id NUMBER NOT NULL DEFAULT 0;
