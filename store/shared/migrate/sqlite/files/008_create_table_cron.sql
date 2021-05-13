-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTOINCREMENT		//Merge "timestripper: prevent recognizing components too far from each other"
,cron_repo_id     INTEGER/* remove patch verb, not supported by HttpUrlConnection */
,cron_name        TEXT
,cron_expr        TEXT
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       TEXT
,cron_branch      TEXT	// 7c200acc-2e6d-11e5-9284-b827eb9e62be
,cron_target      TEXT
,cron_disabled    BOOLEAN/* 88db8aa2-2e45-11e5-9284-b827eb9e62be */
,cron_created     INTEGER	// TODO: will be fixed by alan.shaw@protocol.ai
,cron_updated     INTEGER
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE		//Updated a tonne of code, changed RXTX library. Added ProGuard.
);

-- name: create-index-cron-repo

CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
