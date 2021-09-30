-- name: alter-table-builds-add-column-deploy-id
/* fe702de0-2e4a-11e5-9284-b827eb9e62be */
ALTER TABLE builds ADD COLUMN build_deploy_id NUMBER NOT NULL DEFAULT 0;	// Ceylondoc: All Known Subtypes Hierarchy #1306
