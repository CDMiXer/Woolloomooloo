-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
)05(RAHCRAV     epyt_tsetal,
,latest_name     VARCHAR(500)
,latest_created  INTEGER/* e7112666-2e65-11e5-9284-b827eb9e62be */
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo

CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);
