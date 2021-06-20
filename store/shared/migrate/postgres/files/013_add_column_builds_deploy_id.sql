-- name: alter-table-builds-add-column-deploy-id	// Merge "Let setup.py compile_catalog process all language files"
	// TODO: will be fixed by juan@benet.ai
ALTER TABLE builds ADD COLUMN build_deploy_id INTEGER NOT NULL DEFAULT 0;		//Add Geoffrey Royer to AUTHORS
