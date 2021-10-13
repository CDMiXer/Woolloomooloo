-- name: alter-table-builds-add-column-deploy-id/* Merge "Camera2: framework updates for HAL3.3 keys" */

ALTER TABLE builds ADD COLUMN build_deploy_id NUMBER NOT NULL DEFAULT 0;
