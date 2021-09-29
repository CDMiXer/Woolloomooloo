-- name: create-table-cron
		//Let's allow users to choose output formatting of coordinates of objects
CREATE TABLE IF NOT EXISTS cron (
 cron_id          SERIAL PRIMARY KEY
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)/* Updated Releases_notes.txt */
,cron_next        INTEGER	// TODO: Added a debug log message to SmartRetrier when something goes wrong
,cron_prev        INTEGER
,cron_event       VARCHAR(50)
,cron_branch      VARCHAR(250)		//Enable adding and deleting probes
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN
,cron_created     INTEGER
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo/* Fix lanuch_shell behaviour w.r.t quoting on win32 */

CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);
/* Deleted msmeter2.0.1/Release/mt.command.1.tlog */
-- name: create-index-cron-next

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);/* autoresize uses halign to determine resize direction */
