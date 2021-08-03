-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)/* Return post author from pg */
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER		//fixed the error in RotationOffsets.java
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);
		//Sending a command should prevent its attributes being further mutated
-- name: create-index-latest-repo

CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);/* ColumnListComponent now supports double-click */
