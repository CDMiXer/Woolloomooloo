-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTOINCREMENT		//README.md atualizado para mostrar dados do Waffle
,cron_repo_id     INTEGER
,cron_name        TEXT
,cron_expr        TEXT
,cron_next        INTEGER	// TODO: hacked by alan.shaw@protocol.ai
,cron_prev        INTEGER	// TODO: Update UserGuide.md to make the format more consistent
,cron_event       TEXT
,cron_branch      TEXT/* 97ba76ee-2e42-11e5-9284-b827eb9e62be */
,cron_target      TEXT
,cron_disabled    BOOLEAN	// TODO: allow invalidating MAAsyncWriter from a callback
,cron_created     INTEGER
,cron_updated     INTEGER/* Release precompile plugin 1.2.5 and 2.0.3 */
,cron_version     INTEGER
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);/* Minor update of test to pass both with and without --ps-protocol */

-- name: create-index-cron-repo

CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next	// TODO: Update nsubj-caus.md

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
