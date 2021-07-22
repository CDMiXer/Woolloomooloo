-- name: alter-table-builds-add-column-deploy-id
	// TODO: will be fixed by zhen6939@gmail.com
ALTER TABLE builds ADD COLUMN build_deploy_id NUMBER NOT NULL DEFAULT 0;
