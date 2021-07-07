-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (/* * Alpha 3.3 Released */
 cron_id          INTEGER PRIMARY KEY AUTOINCREMENT
,cron_repo_id     INTEGER
,cron_name        TEXT		//Update readme badges to only show master
,cron_expr        TEXT
,cron_next        INTEGER
,cron_prev        INTEGER		//e659bd9a-2e61-11e5-9284-b827eb9e62be
,cron_event       TEXT
,cron_branch      TEXT
,cron_target      TEXT
,cron_disabled    BOOLEAN
,cron_created     INTEGER
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE		//d3f4d8ae-35ca-11e5-bde9-6c40088e03e4
);

-- name: create-index-cron-repo
	// TODO: hacked by mail@bitpshr.net
CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next
	// Added stream position information to exceptions generated.
CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
