-- name: create-table-latest
/* Fix Lircmap.xml for XBMC */
CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER/* Mark as 0.3.0 Release */
,latest_build_id INTEGER
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)		//Removed jQuery.js as it's no longer required
,latest_created  INTEGER	// TODO: Shader names are fixed
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo

CREATE INDEX ix_latest_repo ON latest (latest_repo_id);	// TODO: - update of Settings access
