-- name: create-table-latest	// TODO: will be fixed by hello@brooklynzelenka.com

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     TEXT -- branch | tag     | pull_request | promote
,latest_name     TEXT -- master | v1.0.0, | 42           | production
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)/* Release of eeacms/forests-frontend:2.0-beta.64 */
);
		//removed the ability to add media to player notes
-- name: create-index-latest-repo
		//75dcf60e-2e4d-11e5-9284-b827eb9e62be
CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);
