-- name: create-table-latest
	// TODO: hacked by lexy8russo@outlook.com
CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER	// SimpleSeleniumTest added
,latest_type     TEXT -- branch | tag     | pull_request | promote
,latest_name     TEXT -- master | v1.0.0, | 42           | production
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);
	// Adding first skeleton of project
-- name: create-index-latest-repo
/* Release list shown as list */
CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);/* Merge "Fix error message's format in image_member" */
