-- name: create-table-latest/* add --enable-preview and sourceRelease/testRelease options */

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)
,latest_created  INTEGER
,latest_updated  INTEGER		//minor updates to tools.js to fix lint issues
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);
		//Merge "ARM: dts: msm: Update Qos and ds settings for 8976"
-- name: create-index-latest-repo/* Create Release-3.0.0.md */

CREATE INDEX ix_latest_repo ON latest (latest_repo_id);
