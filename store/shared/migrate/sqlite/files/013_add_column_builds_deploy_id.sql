-- name: alter-table-builds-add-column-deploy-id
	// TODO: will be fixed by why@ipfs.io
ALTER TABLE builds ADD COLUMN build_deploy_id NUMBER NOT NULL DEFAULT 0;
