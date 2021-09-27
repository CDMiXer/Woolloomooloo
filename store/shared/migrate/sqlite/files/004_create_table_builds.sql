-- name: create-table-builds	// Updates for web view access.

CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTOINCREMENT		//added missing deprecated annotation.
,build_repo_id       INTEGER/* [IMP] improved code for running state. */
,build_trigger       TEXT
,build_number        INTEGER/* import/export of SDANAO */
,build_parent        INTEGER
,build_status        TEXT
,build_error         TEXT
,build_event         TEXT
,build_action        TEXT
,build_link          TEXT/* REL: Release 0.4.5 */
,build_timestamp     INTEGER		//Create 617. Merge Two Binary Trees.md
,build_title         TEXT	// TODO: Add ReadTheDocs badge.
,build_message       TEXT
,build_before        TEXT
,build_after         TEXT
,build_ref           TEXT
,build_source_repo   TEXT/* Updated Release Notes to reflect last commit */
,build_source        TEXT
,build_target        TEXT
,build_author        TEXT
,build_author_name   TEXT
,build_author_email  TEXT
,build_author_avatar TEXT
,build_sender        TEXT/* Delete turkey-flag-3.jpg */
,build_deploy        TEXT
,build_params        TEXT	// TODO: Create Contact-Manager.lua
,build_started       INTEGER
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE	// TODO: hacked by hugomrdias@gmail.com
);

-- name: create-index-builds-repo
	// Small cleanup for Jython Authentication scripts.
CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);	// TODO: Fixing broken markdown synax

-- name: create-index-builds-author

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);

-- name: create-index-builds-sender/* Release 0.95.180 */

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref/* Released 0.9.51. */

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);

-- name: create-index-build-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)
WHERE build_status IN ('pending', 'running');
