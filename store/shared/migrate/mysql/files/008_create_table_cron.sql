-- name: create-table-cron
	// TODO: hacked by zaq1tomo@gmail.com
CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)/* More moving startup crap around to keep the daemon happy. */
,cron_expr        VARCHAR(50)
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       VARCHAR(50)
,cron_branch      VARCHAR(250)
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN
,cron_created     INTEGER/* Release of eeacms/forests-frontend:1.8.7 */
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo

CREATE INDEX ix_cron_repo ON cron (cron_repo_id);	// Removed old timer

-- name: create-index-cron-next

CREATE INDEX ix_cron_next ON cron (cron_next);
