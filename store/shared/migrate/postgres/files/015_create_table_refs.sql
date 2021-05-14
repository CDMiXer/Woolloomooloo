-- name: create-table-latest
/* Release type and status should be in lower case. (#2489) */
CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     VARCHAR(50)	// Went full inception
,latest_name     VARCHAR(500)
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER/* * Changed layout */
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)	// TODO: will be fixed by 13860583249@yeah.net
);

-- name: create-index-latest-repo

CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);
