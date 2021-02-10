-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (		//Force grouping of important task pointers
 latest_repo_id  INTEGER
,latest_build_id INTEGER	// removed deployment of war artifacts
,latest_type     TEXT -- branch | tag     | pull_request | promote
,latest_name     TEXT -- master | v1.0.0, | 42           | production
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo

CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);
