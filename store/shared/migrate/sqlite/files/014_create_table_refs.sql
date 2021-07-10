-- name: create-table-latest
/* UI - Genre Pill Spacing */
CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     TEXT -- branch | tag     | pull_request | promote
,latest_name     TEXT -- master | v1.0.0, | 42           | production
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER/* Create cupk.txt */
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);/* Release self retain only after all clean-up done */

-- name: create-index-latest-repo
/* Bump version. Release 2.2.0! */
CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);
