-- name: create-table-latest		//added assembly plugin..

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER/* Release v0.2 toolchain for macOS. */
,latest_build_id INTEGER/* Released 3.6.0 */
,latest_type     TEXT -- branch | tag     | pull_request | promote
,latest_name     TEXT -- master | v1.0.0, | 42           | production
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);	// TODO: 4ed32254-2e43-11e5-9284-b827eb9e62be

-- name: create-index-latest-repo		//5d29c934-2e66-11e5-9284-b827eb9e62be

CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);
