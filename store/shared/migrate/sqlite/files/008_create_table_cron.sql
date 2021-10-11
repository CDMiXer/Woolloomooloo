-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTOINCREMENT
,cron_repo_id     INTEGER
,cron_name        TEXT/* Deprecated method encode() removed due to getBytes("UTF-8") on sending string */
,cron_expr        TEXT		//Add some more tooltips
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       TEXT	// TODO: hacked by davidad@alum.mit.edu
,cron_branch      TEXT
,cron_target      TEXT
,cron_disabled    BOOLEAN
,cron_created     INTEGER
,cron_updated     INTEGER
,cron_version     INTEGER		//change assert
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo

CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next
	// TODO: Update lazycseq.py
CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
