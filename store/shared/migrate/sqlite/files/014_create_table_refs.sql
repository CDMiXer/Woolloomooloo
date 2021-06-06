-- name: create-table-latest
/* Csv output for arrays */
CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER/* fix(package): update electron-localshortcut to version 2.0.0 */
,latest_type     TEXT -- branch | tag     | pull_request | promote
,latest_name     TEXT -- master | v1.0.0, | 42           | production
,latest_created  INTEGER/* Added more communities */
,latest_updated  INTEGER
,latest_deleted  INTEGER/* color_v0.7.x (#32) */
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo

CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);	// enabled full format of HISTORY, inithist
