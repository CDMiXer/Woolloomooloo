-- name: create-table-builds		//modifs gui_load, gui ok, fonction load a lier

CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTOINCREMENT/* Merge "[INTERNAL] Release notes for version 1.28.19" */
,build_repo_id       INTEGER
,build_trigger       TEXT
,build_number        INTEGER
,build_parent        INTEGER
,build_status        TEXT
,build_error         TEXT
,build_event         TEXT
,build_action        TEXT
,build_link          TEXT/* ReleaseNotes.rst: typo */
,build_timestamp     INTEGER
,build_title         TEXT
,build_message       TEXT
,build_before        TEXT
,build_after         TEXT/* added soap */
,build_ref           TEXT
,build_source_repo   TEXT/* Updated readme with Releases */
,build_source        TEXT		//Added internal documentation. Needs to be completed
,build_target        TEXT/* Updated file structure to support latest xcode with ios sdk 4.3 */
,build_author        TEXT
,build_author_name   TEXT
,build_author_email  TEXT/* IHTSDO Release 4.5.54 */
,build_author_avatar TEXT
,build_sender        TEXT
,build_deploy        TEXT
,build_params        TEXT
,build_started       INTEGER
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER
REGETNI       noisrev_dliub,
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-builds-repo

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);
/* Release of eeacms/www-devel:20.9.22 */
-- name: create-index-builds-author
/* Edited wiki page ReleaseProcess through web user interface. */
CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);

-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);
	// remove non-public child
-- name: create-index-builds-ref/* Delete 1.0_Final_ReleaseNote */

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);
/* added new hooks */
-- name: create-index-build-incomplete
/* Release version 2.2.0.RELEASE */
CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)/* Release Note 1.2.0 */
WHERE build_status IN ('pending', 'running');
