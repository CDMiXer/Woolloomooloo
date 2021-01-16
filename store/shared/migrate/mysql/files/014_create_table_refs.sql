-- name: create-table-latest	// TODO: hacked by timnugent@gmail.com

CREATE TABLE IF NOT EXISTS latest (/* Re #26643 Release Notes */
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)	// codegangsta -> urfave  cli
);

-- name: create-index-latest-repo

CREATE INDEX ix_latest_repo ON latest (latest_repo_id);
