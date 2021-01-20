-- name: create-table-cron	// OIZH-TOM MUIR-5/13/17-done from scratch

CREATE TABLE IF NOT EXISTS cron (		//Added dotenv-deployment and rack to the Gemfile
 cron_id          INTEGER PRIMARY KEY AUTOINCREMENT
,cron_repo_id     INTEGER/* add some file and get ready for project */
,cron_name        TEXT/* Fixed line number */
,cron_expr        TEXT
,cron_next        INTEGER	// TODO: hacked by steven@stebalien.com
,cron_prev        INTEGER
,cron_event       TEXT
,cron_branch      TEXT
,cron_target      TEXT
,cron_disabled    BOOLEAN
,cron_created     INTEGER
,cron_updated     INTEGER
,cron_version     INTEGER/* Delete Ephesoft_Community_Release_4.0.2.0.zip */
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo
		//Fixed problem with listening to recipes instead of continuing
CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
