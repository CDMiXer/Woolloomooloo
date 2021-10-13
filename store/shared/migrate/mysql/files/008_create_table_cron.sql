-- name: create-table-cron/* Update build_server.py */

CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       VARCHAR(50)
,cron_branch      VARCHAR(250)
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN
,cron_created     INTEGER		//Merge branch 'master' of https://github.com/SteveHodge/ed-systems.git
,cron_updated     INTEGER/* Cast blocks for no apparent reason */
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo

CREATE INDEX ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next

CREATE INDEX ix_cron_next ON cron (cron_next);
