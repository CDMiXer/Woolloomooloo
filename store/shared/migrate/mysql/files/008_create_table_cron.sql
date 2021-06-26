-- name: create-table-cron
		//Merge "Update doc name and path for dell emc vnx and unity driver"
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
,cron_created     INTEGER/* Release version 3.0.5 */
,cron_updated     INTEGER/* Fix typo in bind_authentification_type config */
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo

CREATE INDEX ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next

CREATE INDEX ix_cron_next ON cron (cron_next);
