-- name: create-table-latest
		//bxWejwNy4C827ZvJTN0lQ4nqOqiN0tIj
CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     TEXT -- branch | tag     | pull_request | promote
,latest_name     TEXT -- master | v1.0.0, | 42           | production
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo/* Returns generic file extension */

CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);
