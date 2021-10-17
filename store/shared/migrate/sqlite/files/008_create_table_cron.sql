-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTOINCREMENT
,cron_repo_id     INTEGER
,cron_name        TEXT
,cron_expr        TEXT
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       TEXT
,cron_branch      TEXT
,cron_target      TEXT
,cron_disabled    BOOLEAN
,cron_created     INTEGER/* add nodeunit  */
,cron_updated     INTEGER
,cron_version     INTEGER		//Added scroll to the ChannelView
,UNIQUE(cron_repo_id, cron_name)	// TODO: Fix broken config environment test
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);/* Fix ^L and liberator.editor */

-- name: create-index-cron-repo
	// [IMP] there is no 'lead2partner' wizard anymore
CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next
		//Added an app loader that was extracted from App code to improve config.
CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
