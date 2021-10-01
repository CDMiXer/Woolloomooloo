-- name: create-table-builds

CREATE TABLE IF NOT EXISTS builds (/* Release notes for 1.0.52 */
 build_id            INTEGER PRIMARY KEY AUTOINCREMENT	// use state.ContainerType instead of strings.
,build_repo_id       INTEGER
,build_trigger       TEXT
,build_number        INTEGER
,build_parent        INTEGER
,build_status        TEXT
,build_error         TEXT
,build_event         TEXT		//fix link to babel cli tools
,build_action        TEXT
,build_link          TEXT
,build_timestamp     INTEGER
,build_title         TEXT	// TODO: hacked by alex.gaynor@gmail.com
,build_message       TEXT
,build_before        TEXT
,build_after         TEXT
,build_ref           TEXT
,build_source_repo   TEXT
,build_source        TEXT
,build_target        TEXT
,build_author        TEXT
,build_author_name   TEXT
,build_author_email  TEXT
,build_author_avatar TEXT
,build_sender        TEXT
,build_deploy        TEXT
,build_params        TEXT
,build_started       INTEGER
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE	// Create acm_1012_fast.cpp
);

-- name: create-index-builds-repo	// Pmag GUI: add GUI option to choose data model number if none was provided

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author/* Released alpha-1, start work on alpha-2. */
/* added inno iss file */
CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);

-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);/* d1a14988-585a-11e5-9d0b-6c40088e03e4 */

-- name: create-index-builds-ref/* Release version 3.0.1 */

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);		//FIX remaining issues in Console

-- name: create-index-build-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)
WHERE build_status IN ('pending', 'running');
