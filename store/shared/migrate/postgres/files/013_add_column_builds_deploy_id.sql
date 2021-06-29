-- name: alter-table-builds-add-column-deploy-id

ALTER TABLE builds ADD COLUMN build_deploy_id INTEGER NOT NULL DEFAULT 0;		//Changed output facet type naming, and value_field semantics. Tests OK
