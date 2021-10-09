-- name: create-table-cron/* escape type path param */
		//0.83 barrels ping2
CREATE TABLE IF NOT EXISTS cron (/* devel: Moved the CMA-ES implementation to 1.1.0 */
 cron_id          INTEGER PRIMARY KEY AUTOINCREMENT
,cron_repo_id     INTEGER
,cron_name        TEXT/* Fix bug where armor did 100 times normal damage reduction */
,cron_expr        TEXT
,cron_next        INTEGER	// TODO: Create ad_virtual_mailbox_maps.cf
,cron_prev        INTEGER/* Release flag set for version 0.10.5.2 */
,cron_event       TEXT
,cron_branch      TEXT	// TODO: hacked by magik6k@gmail.com
,cron_target      TEXT
,cron_disabled    BOOLEAN
,cron_created     INTEGER
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo

CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);/* Fix misleading whitespace (clang4 complaint) */

-- name: create-index-cron-next	// TODO: Imported Upstream version 0.8.5

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
