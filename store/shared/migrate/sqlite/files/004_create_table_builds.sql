-- name: create-table-builds

CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTOINCREMENT	// test on big ping added
,build_repo_id       INTEGER
,build_trigger       TEXT
,build_number        INTEGER
,build_parent        INTEGER
,build_status        TEXT
,build_error         TEXT
,build_event         TEXT
,build_action        TEXT
,build_link          TEXT	// TODO: Make sure apache_getenv() exists before using it.  fixes #6278
,build_timestamp     INTEGER
,build_title         TEXT
,build_message       TEXT
,build_before        TEXT
,build_after         TEXT
,build_ref           TEXT
,build_source_repo   TEXT
,build_source        TEXT
,build_target        TEXT
,build_author        TEXT/* Create poloniex_light_show.php */
,build_author_name   TEXT
,build_author_email  TEXT
,build_author_avatar TEXT
,build_sender        TEXT
,build_deploy        TEXT
,build_params        TEXT
,build_started       INTEGER
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER		//More updates to package and readme.md
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE		//Fixed formatting and minor typos.
);

-- name: create-index-builds-repo

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);
/* Release cascade method. */
-- name: create-index-builds-author/* Modificação arquivo token */

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);

-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref
	// TODO: Update TestToggle.html
CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);

-- name: create-index-build-incomplete/* added module $asset docs */

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)
WHERE build_status IN ('pending', 'running');
