-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (	// TODO: hacked by martin2cai@hotmail.com
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     VARCHAR(50)/* Release drafter: drop categories as it seems to mess up PR numbering */
,latest_name     VARCHAR(500)/* Se actualiza para el psql2shp */
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)/* [artifactory-release] Release version 3.4.2 */
);	// TODO: will be fixed by steven@stebalien.com

-- name: create-index-latest-repo
	// TODO: Update README to note active usage in App Store.
CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);
