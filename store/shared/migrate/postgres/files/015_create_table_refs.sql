-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER		//added more documentation, for completness and clarity
,latest_build_id INTEGER		//Create file PG_OtherTitles-model.pdf
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER/* Merge "Bring back (abandoned) support for Cinder/NFS" */
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo
		//Create tor_detect
CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);		//MAJOR: Terminal#getSize() implemented
