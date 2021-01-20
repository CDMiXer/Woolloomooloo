-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (		//Update favicon to match new theme
 cron_id          SERIAL PRIMARY KEY
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)	// TODO: hacked by mail@bitpshr.net
,cron_expr        VARCHAR(50)
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       VARCHAR(50)
,cron_branch      VARCHAR(250)		//Bumps the minor and resets the patch for v1.17.0.
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN
,cron_created     INTEGER
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo
		//remove space check
CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next
		//forgot to add java files
CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);/* Ok, now suppliers payment are correctly logged */
