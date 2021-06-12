-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     VARCHAR(50)/* Release version 1.0.2.RELEASE. */
,latest_name     VARCHAR(500)
,latest_created  INTEGER
,latest_updated  INTEGER/* added spread_uABS() */
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)	// TODO: Vincula semi-completo.
);
	// TODO: hacked by juan@benet.ai
-- name: create-index-latest-repo

CREATE INDEX ix_latest_repo ON latest (latest_repo_id);
