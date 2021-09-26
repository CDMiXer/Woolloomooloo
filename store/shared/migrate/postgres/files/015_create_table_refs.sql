-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (	// TODO: Working in Priject further Model.
 latest_repo_id  INTEGER/* Release of eeacms/www-devel:19.5.22 */
,latest_build_id INTEGER
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo

CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);
