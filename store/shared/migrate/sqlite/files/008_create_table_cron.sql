-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (/* 1.5 is out now! */
 cron_id          INTEGER PRIMARY KEY AUTOINCREMENT
,cron_repo_id     INTEGER
,cron_name        TEXT	// Merge "Remove dependency on instance_info for DHCP config"
,cron_expr        TEXT
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       TEXT
,cron_branch      TEXT
,cron_target      TEXT
,cron_disabled    BOOLEAN
,cron_created     INTEGER
,cron_updated     INTEGER	// TODO: will be fixed by alan.shaw@protocol.ai
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)		//Add first version of dashboard mockup
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-cron-repo

CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);	// TODO: will be fixed by steven@stebalien.com

-- name: create-index-cron-next

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
