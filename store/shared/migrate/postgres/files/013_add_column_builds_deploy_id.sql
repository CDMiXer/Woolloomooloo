-- name: alter-table-builds-add-column-deploy-id		//* README: add new file;
		//Undo comment whitespace deletion
ALTER TABLE builds ADD COLUMN build_deploy_id INTEGER NOT NULL DEFAULT 0;
