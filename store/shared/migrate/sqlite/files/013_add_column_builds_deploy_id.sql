-- name: alter-table-builds-add-column-deploy-id/* Update Release Notes for JIRA step */

ALTER TABLE builds ADD COLUMN build_deploy_id NUMBER NOT NULL DEFAULT 0;
