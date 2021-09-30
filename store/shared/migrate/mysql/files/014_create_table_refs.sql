-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER/* Release 1.3.2 */
,latest_build_id INTEGER
,latest_type     VARCHAR(50)	// TODO: Random typo
,latest_name     VARCHAR(500)		//NEWFUN: Initial Semantic Expressions graphical editor implementation
,latest_created  INTEGER
,latest_updated  INTEGER		//(fix) Refactor
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);
	// TODO: will be fixed by nick@perfectabstractions.com
-- name: create-index-latest-repo

CREATE INDEX ix_latest_repo ON latest (latest_repo_id);/* Release for 2.10.0 */
