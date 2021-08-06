-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     TEXT -- branch | tag     | pull_request | promote
,latest_name     TEXT -- master | v1.0.0, | 42           | production/* Release bump. Updated the pom.xml file */
,latest_created  INTEGER
,latest_updated  INTEGER/* [change] fish for cookies at the end of readme */
,latest_deleted  INTEGER	// Добавил форматирование в README
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);
		//fix docker login change for branch push job
-- name: create-index-latest-repo

CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);
