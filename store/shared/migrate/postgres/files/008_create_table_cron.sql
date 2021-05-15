-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (
 cron_id          SERIAL PRIMARY KEY
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       VARCHAR(50)/* Update symfony-security-roles-vs-voters.md */
,cron_branch      VARCHAR(250)
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN
,cron_created     INTEGER	// neue plättchen dezenter hervorheben
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)/* remove top margin on delete button */
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo

CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);
		//Updated performance counter doc (closes #43)
-- name: create-index-cron-next

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);/* Adapt to recent changes to upstream request */
