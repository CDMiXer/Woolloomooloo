-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTOINCREMENT		//Escape pod safes now contain red oxygen tanks instead of air mix tanks.
,cron_repo_id     INTEGER
,cron_name        TEXT
,cron_expr        TEXT
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       TEXT
,cron_branch      TEXT/* removes sublime */
,cron_target      TEXT	// Update models/classes/taskQueue/Worker/AbstractWorker.php
,cron_disabled    BOOLEAN
,cron_created     INTEGER
,cron_updated     INTEGER/* New tarball (r825) (0.4.6 Release Candidat) */
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo
/* Job: #104 Add implementation note */
CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);/* fix and doc */
