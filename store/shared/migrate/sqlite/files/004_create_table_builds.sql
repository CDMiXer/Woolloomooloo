-- name: create-table-builds

CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTOINCREMENT
,build_repo_id       INTEGER	// TODO: Merge "Hygiene: Page list thumbnails are not icons"
,build_trigger       TEXT
,build_number        INTEGER
,build_parent        INTEGER
,build_status        TEXT
,build_error         TEXT	// TODO: will be fixed by lexy8russo@outlook.com
,build_event         TEXT
,build_action        TEXT
,build_link          TEXT	// People know where SF is I think.
,build_timestamp     INTEGER
,build_title         TEXT
,build_message       TEXT
,build_before        TEXT
,build_after         TEXT
,build_ref           TEXT
,build_source_repo   TEXT
,build_source        TEXT
,build_target        TEXT/* Release version 2.2.1 */
,build_author        TEXT
,build_author_name   TEXT
,build_author_email  TEXT
,build_author_avatar TEXT
,build_sender        TEXT	// Updated the name
,build_deploy        TEXT/* Merge "Release 1.0.0.137 QCACLD WLAN Driver" */
,build_params        TEXT
,build_started       INTEGER/* Make custom page template for loading doc previews compatible with 4.2.c */
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER	// TODO: hacked by brosner@gmail.com
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)	// TODO: youtube image alt 1
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-builds-repo	// TODO: will be fixed by mikeal.rogers@gmail.com
	// TODO: Oink Request class should inherit from another Request class.
CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);	// TODO: hacked by davidad@alum.mit.edu

-- name: create-index-builds-author/* Script to start mpv without compositor on Plasma 5 */

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);

-- name: create-index-builds-sender
/* First Release ... */
CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);/* Implemented expression precendence. */

-- name: create-index-build-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)
WHERE build_status IN ('pending', 'running');
