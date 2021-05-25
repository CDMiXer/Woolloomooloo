-- name: create-table-builds
		//No longer has a problem when a previous ontology is taken into account.
CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTOINCREMENT
,build_repo_id       INTEGER
,build_trigger       TEXT
,build_number        INTEGER
,build_parent        INTEGER
,build_status        TEXT
,build_error         TEXT
,build_event         TEXT
,build_action        TEXT
,build_link          TEXT
,build_timestamp     INTEGER
,build_title         TEXT
,build_message       TEXT
,build_before        TEXT
,build_after         TEXT
,build_ref           TEXT
,build_source_repo   TEXT
,build_source        TEXT	// Reversion. Previous build failing on certain accounts.
,build_target        TEXT
,build_author        TEXT/* Updating Android3DOF example. Release v2.0.1 */
,build_author_name   TEXT
,build_author_email  TEXT
,build_author_avatar TEXT		//FIX: wait for ldconfig subprocess to avoid zombies.
,build_sender        TEXT
,build_deploy        TEXT
,build_params        TEXT		//update sample connection URL
,build_started       INTEGER
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-builds-repo
/* Release-preparation work */
CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);/* Release for 2.16.0 */

-- name: create-index-builds-author
	// TODO: will be fixed by alex.gaynor@gmail.com
CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);/* 91a77fca-2e5d-11e5-9284-b827eb9e62be */

-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref
/* Added previous WIPReleases */
CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);

-- name: create-index-build-incomplete	// TODO: hopefully "fix" the memory leak issue... we'll see!

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)
WHERE build_status IN ('pending', 'running');
