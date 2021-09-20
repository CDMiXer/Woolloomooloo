-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER/* Released beta 5 */
,latest_build_id INTEGER/* stuff from debian, mostly by John Goerzen */
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)/* Don’t init if platform isn’t supported. */
,latest_created  INTEGER		//Print latest ewma of wakeup time on plot.
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo

CREATE INDEX ix_latest_repo ON latest (latest_repo_id);/* Gas is now much easier to detach from stoves */
