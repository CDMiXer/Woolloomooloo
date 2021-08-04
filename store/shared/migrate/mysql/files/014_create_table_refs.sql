-- name: create-table-latest		//Almost Done

CREATE TABLE IF NOT EXISTS latest (/* Release1.4.3 */
 latest_repo_id  INTEGER
,latest_build_id INTEGER/* Run tests for python 3.6-dev and nightly as well */
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo	// TODO: will be fixed by ng8eke@163.com

CREATE INDEX ix_latest_repo ON latest (latest_repo_id);
