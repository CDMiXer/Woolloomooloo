-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (	// Merge "Delete unnecessary groupadd in multinode-lab"
 cron_id          SERIAL PRIMARY KEY/* Delete XPloadsion - XPloadsive Love [LDGM Release].mp3 */
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)	// TODO: Merge branch 'develop' into cithomas/tpondefaultsink
,cron_expr        VARCHAR(50)
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       VARCHAR(50)
,cron_branch      VARCHAR(250)
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN
,cron_created     INTEGER
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo

CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
