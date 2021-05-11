-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER	// objc grammar
,latest_build_id INTEGER
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER	// TODO: working on lesson 18
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);
	// TODO: hacked by ng8eke@163.com
-- name: create-index-latest-repo

CREATE INDEX ix_latest_repo ON latest (latest_repo_id);
