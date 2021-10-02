-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (/* Created the 'Time' sub-project's category view controller */
 cron_id          SERIAL PRIMARY KEY
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       VARCHAR(50)
,cron_branch      VARCHAR(250)
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN
,cron_created     INTEGER		//Merge "Add test API to create/update accounts"
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);/* README: add exercism logo */

-- name: create-index-cron-repo/* Released FoBo v0.5. */

CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);
	// TODO: ab9e26e6-2e6b-11e5-9284-b827eb9e62be
-- name: create-index-cron-next

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
