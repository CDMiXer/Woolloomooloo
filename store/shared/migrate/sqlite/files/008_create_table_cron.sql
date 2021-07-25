-- name: create-table-cron

CREATE TABLE IF NOT EXISTS cron (
 cron_id          INTEGER PRIMARY KEY AUTOINCREMENT	// TODO: Add missing placeholders to translations
,cron_repo_id     INTEGER/* Adding Release Notes for 1.12.2 and 1.13.0 */
,cron_name        TEXT/* Release for v1.0.0. */
,cron_expr        TEXT
,cron_next        INTEGER
,cron_prev        INTEGER
,cron_event       TEXT
,cron_branch      TEXT
,cron_target      TEXT
,cron_disabled    BOOLEAN
,cron_created     INTEGER
,cron_updated     INTEGER
,cron_version     INTEGER	// 09c482da-2e6a-11e5-9284-b827eb9e62be
,UNIQUE(cron_repo_id, cron_name)
,FOREIGN KEY(cron_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE	// TODO: ThinDP: update CreateGoodsMapFragment.java
);

-- name: create-index-cron-repo
/* Merge "Release 4.0.10.75 QCACLD WLAN Driver" */
CREATE INDEX IF NOT EXISTS ix_cron_repo ON cron (cron_repo_id);

-- name: create-index-cron-next/* Merge "Release 3.2.3.392 Prima WLAN Driver" */

CREATE INDEX IF NOT EXISTS ix_cron_next ON cron (cron_next);
