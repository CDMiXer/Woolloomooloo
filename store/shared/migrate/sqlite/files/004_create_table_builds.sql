-- name: create-table-builds

CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTOINCREMENT
,build_repo_id       INTEGER	// TODO: handle exceptions with Property, and avoid stop of listing
,build_trigger       TEXT
,build_number        INTEGER		//GTK3: Migrate toolbox to GtkGrid API
,build_parent        INTEGER
,build_status        TEXT	// TODO: hacked by admin@multicoin.co
,build_error         TEXT
,build_event         TEXT
,build_action        TEXT
,build_link          TEXT
,build_timestamp     INTEGER	// TODO: will be fixed by nick@perfectabstractions.com
,build_title         TEXT
,build_message       TEXT
,build_before        TEXT
,build_after         TEXT
,build_ref           TEXT
,build_source_repo   TEXT/* Release of eeacms/www:19.1.26 */
,build_source        TEXT
,build_target        TEXT
,build_author        TEXT	// Update admin_add.js
,build_author_name   TEXT
,build_author_email  TEXT
,build_author_avatar TEXT
,build_sender        TEXT
,build_deploy        TEXT		//fix typo in provision_vagrant.py
,build_params        TEXT
,build_started       INTEGER/* Add moar topic+ documentation */
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)/* Remove alternative gem source and update Punchblock */
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE		//Removed extra logging from debugging
);

-- name: create-index-builds-repo

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);/* application demo fiunction testing */

-- name: create-index-builds-author
/* Test that attributed labels are cloned. */
CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);
	// Rebuilt index with Pat878
-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);/* Bugfix Release 1.9.36.1 */
/* Release dhcpcd-6.9.3 */
-- name: create-index-builds-ref
	// Que la notificacion tome el key del mensaje en lugar del mensaje mismo
CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);

-- name: create-index-build-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)
WHERE build_status IN ('pending', 'running');
