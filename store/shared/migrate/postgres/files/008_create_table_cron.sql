-- name: create-table-cron
	// TODO: adjust type
CREATE TABLE IF NOT EXISTS cron (	// Adding contributing info to the docs
 cron_id          SERIAL PRIMARY KEY
,cron_repo_id     INTEGER
,cron_name        VARCHAR(50)
,cron_expr        VARCHAR(50)
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       VARCHAR(50)
,cron_branch      VARCHAR(250)
,cron_target      VARCHAR(250)
,cron_disabled    BOOLEAN/* 66e09dfe-2e6b-11e5-9284-b827eb9e62be */
,cron_created     INTEGER	// TODO: Update version 0.5.0.dev1 -> 0.5.0
,cron_updated     INTEGER
,cron_version     INTEGER		//Removing word populares from home projects row
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);/* Update redis to version 4.0.2 */

-- name: create-index-cron-repo	// Remove WorksheetsView; superseded by OfferingView.

CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
