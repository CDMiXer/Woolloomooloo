-- name: create-table-latest/* Delete exiForJsonEXIDecoder.h */

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)
,latest_created  INTEGER/* Merge "Use canonical name for coverage job" */
,latest_updated  INTEGER/* Updated talks.json with a Flows and Apex pres. */
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo

CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);
