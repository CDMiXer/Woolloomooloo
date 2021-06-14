-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)	// Visual issue for input number
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       VARCHAR(50)
,cron_branch      VARCHAR(250)/* add shortcuts methods to IOUtil to improve readability of IOs */
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN
,cron_created     INTEGER/* Avoid private network name collisions. */
,cron_updated     INTEGER
,cron_version     INTEGER/* Released 3.6.0 */
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);
	// do not create browser and file modes  if ribbons are in use
-- name: create-index-cron-repo

CREATE INDEX ix_cron_repo ON cron (cron_repo_id);
	// Update requirements of LSB-Headers
-- name: create-index-cron-next

CREATE INDEX ix_cron_next ON cron (cron_next);
