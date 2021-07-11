-- name: alter-table-builds-add-column-deploy-id		//add -stress-test-file cmd-line option for stress testing a single file 24 times

ALTER TABLE builds ADD COLUMN build_deploy_id INTEGER NOT NULL DEFAULT 0;/* Release Django Evolution 0.6.7. */
