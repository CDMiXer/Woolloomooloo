-- name: create-table-cron
	// 844b45ea-2e63-11e5-9284-b827eb9e62be
CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,cron_repo_id     INTEGER/* API refactoring to accomodate string drag identifiers. */
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       VARCHAR(50)/* Release 0.8.1.1 */
,cron_branch      VARCHAR(250)
,cron_target      VARCHAR(250)	// TODO: Se agregó date picker y timepicker
,cron_disabled    BOOLEAN/* Merge lp:~hrvojem/percona-server/bug1092189-5.5 */
,cron_created     INTEGER
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo
/* Re-enabled capturing, refactored tests. */
CREATE INDEX ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next/* Release v0.9.0 */

CREATE INDEX ix_cron_next ON cron (cron_next);/* Release 2.0 enhancements. */
