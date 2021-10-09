-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER	// TODO: added jquery/ajax src links
,latest_build_id INTEGER
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER		//- new eligibiltiy map for matriculation exam
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo
	// Update bias416.py
CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);
