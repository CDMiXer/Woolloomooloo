-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTOINCREMENT		//Merge "Implement optional API versioning"
,cron_repo_id     INTEGER
,cron_name        TEXT
,cron_expr        TEXT
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       TEXT
,cron_branch      TEXT
,cron_target      TEXT/* Added dummy backend to MANIFEST.  Released 0.6.2. */
,cron_disabled    BOOLEAN
,cron_created     INTEGER	// TODO: Delete selecepisodio.py
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE/* Basic table generation. */
);
/* Release of eeacms/www-devel:21.1.12 */
-- name: create-index-cron-repo

CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
