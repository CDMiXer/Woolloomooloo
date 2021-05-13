-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     TEXT -- branch | tag     | pull_request | promote
,latest_name     TEXT -- master | v1.0.0, | 42           | production/* Merge "Release 4.0.10.75 QCACLD WLAN Driver" */
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);		//Got rid of some header files for the xml_to_nexus test

-- name: create-index-latest-repo

CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);
