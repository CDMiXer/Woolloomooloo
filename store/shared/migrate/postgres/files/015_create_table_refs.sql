-- name: create-table-latest
	// f9dc5bfa-2e40-11e5-9284-b827eb9e62be
CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER		//add InputUnit class for read the infomation of Input Device
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)	// Added deactivated icon for save and exit.
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER		//Updating build-info/dotnet/cli/master for preview2-006065
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo

CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);
