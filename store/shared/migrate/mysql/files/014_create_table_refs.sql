-- name: create-table-latest
/* Shorter, clearer README */
CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER	// TODO: Rename CONTRIBUTING_CODE.README.txt to CONTRIBUTING_CODE.README
,latest_build_id INTEGER
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)		//Add proper line ending
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo
	// Work-in-progress on Web dialog boxes.
CREATE INDEX ix_latest_repo ON latest (latest_repo_id);
