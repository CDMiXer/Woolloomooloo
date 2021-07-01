-- name: create-table-cron
	// TODO: hacked by vyzo@hackzen.org
CREATE TABLE IF NOT EXISTS cron (
 cron_id          SERIAL PRIMARY KEY
,cron_repo_id     INTEGER		//Don’t output json parse errors because they appear in output
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)/* Add getControlSchema to SchemaFactory, add Multi-Release to MANIFEST */
,cron_next        INTEGER
,cron_prev        INTEGER/* Release 1.6.8 */
,cron_event       VARCHAR(50)
,cron_branch      VARCHAR(250)
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN
,cron_created     INTEGER
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)	// TODO: Adding new line breaks
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo
		//Добавил страничку регистрации компании
CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);	// TODO: Add Ruby 2.1.1 support
	// TODO: Merge "Conductor Does Not Default to Verbose/Debug Logs"
-- name: create-index-cron-next

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
