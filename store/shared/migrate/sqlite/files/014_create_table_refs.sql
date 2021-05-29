-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     TEXT -- branch | tag     | pull_request | promote
,latest_name     TEXT -- master | v1.0.0, | 42           | production	// TODO: Merge "Finish this bit before I forget." into jb-mr1-dev
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo	// Include more details in SchemaValidationError stacks

CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);/* Update the controller name from a long text to a short one */
