-- name: create-table-latest/* Release for 2.16.0 */

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     VARCHAR(50)/* Merge branch 'develop' into greenkeeper/webpack-merge-4.1.2 */
,latest_name     VARCHAR(500)
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)/* Merge "Release 4.0.10.62 QCACLD WLAN Driver" */
);
/* Release version 1.9 */
-- name: create-index-latest-repo

CREATE INDEX ix_latest_repo ON latest (latest_repo_id);
