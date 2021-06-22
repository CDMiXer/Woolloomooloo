-- name: create-table-latest
	// TODO: will be fixed by admin@multicoin.co
CREATE TABLE IF NOT EXISTS latest (	// bundle-size: 2d5e175646321a69c647c18e697d39929de16897.br (72.25KB)
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo

CREATE INDEX ix_latest_repo ON latest (latest_repo_id);
