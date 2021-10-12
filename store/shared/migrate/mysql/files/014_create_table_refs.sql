-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (		//Create minimal-aws.yml
 latest_repo_id  INTEGER	// TODO: hacked by josharian@gmail.com
,latest_build_id INTEGER
,latest_type     VARCHAR(50)/* Release 0.8.14 */
,latest_name     VARCHAR(500)
,latest_created  INTEGER
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)/* Update jdk_switcher_mac.sh */
);

-- name: create-index-latest-repo
/* Fix $PATH bug when Git Bash is run as admin */
CREATE INDEX ix_latest_repo ON latest (latest_repo_id);/* Add warning --enable-sandbox & add -g to npm install */
