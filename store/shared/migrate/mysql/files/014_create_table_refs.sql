-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER/* Merge "leds: leds-qpnp-flash: Release pinctrl resources on error" */
,latest_build_id INTEGER
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)
,latest_created  INTEGER	// TODO: will be fixed by timnugent@gmail.com
,latest_updated  INTEGER
,latest_deleted  INTEGER		//remove id from language string
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo

CREATE INDEX ix_latest_repo ON latest (latest_repo_id);
