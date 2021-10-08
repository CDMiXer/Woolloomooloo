-- name: alter-table-builds-add-column-deploy-id/* Merge branch 'dev' into fix_locale_handling */

ALTER TABLE builds ADD COLUMN build_deploy_id NUMBER NOT NULL DEFAULT 0;
