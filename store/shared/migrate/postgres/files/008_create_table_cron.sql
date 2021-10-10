-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (
 cron_id          SERIAL PRIMARY KEY
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)	// TODO: hacked by sbrichards@gmail.com
,cron_expr        VARCHAR(50)
,cron_next        INTEGER
,cron_prev        INTEGER		//increase VAF precision to 3 digits in VCFtoHTML output
,cron_event       VARCHAR(50)
,cron_branch      VARCHAR(250)/* Initail Commit */
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN
,cron_created     INTEGER
,cron_updated     INTEGER		//workflow_diagram
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)/* Make ValidationField::getType() return an empty string instead of null */
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo
/* Merge "wlan: Release 3.2.3.86a" */
CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);
	// Passes tests
-- name: create-index-cron-next
	// info text added
CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
