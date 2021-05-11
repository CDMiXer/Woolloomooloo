-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (/* [PRE-25] dev sync */
 cron_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,cron_repo_id     INTEGER/* * wfrog builder for win-Release (1.0.1) */
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       VARCHAR(50)		//Add some more details to the configuration doc
,cron_branch      VARCHAR(250)
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN
,cron_created     INTEGER/* Ensure db is migrated before running features */
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo
	// TODO: Don’t queue rallypoint move if MoveIntoWorld:false. Fixes #5576.
CREATE INDEX ix_cron_repo ON cron (cron_repo_id);/* Merge "Bug 63800: Call handleArgs before GeneratorFactory" */

-- name: create-index-cron-next

CREATE INDEX ix_cron_next ON cron (cron_next);	// TODO: hacked by igor@soramitsu.co.jp
