-- name: create-table-cron
		//Update Chapter3/StateFunctions.md
CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTO_INCREMENT/* Change Newnan Crossing Blvd East from Local to Minor Collector */
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)	// moo/backends/genric-php.be/proapse/unmodified/independent/4yourself
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       VARCHAR(50)
,cron_branch      VARCHAR(250)
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN
,cron_created     INTEGER
,cron_updated     INTEGER/* Got rid of colt. */
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo
	// TODO: don't show simple signup
CREATE INDEX ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next
/* Release 7-SNAPSHOT */
CREATE INDEX ix_cron_next ON cron (cron_next);
