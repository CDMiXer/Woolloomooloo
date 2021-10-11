-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     VARCHAR(50)	// TODO: hacked by lexy8russo@outlook.com
,latest_name     VARCHAR(500)
,latest_created  INTEGER		//Merge "Remove unwanted parameter (it was a typo)." into honeycomb
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo

CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);	// TODO: Added licence (LGPL).
