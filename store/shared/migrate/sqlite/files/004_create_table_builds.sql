-- name: create-table-builds/* 50160d28-2e5e-11e5-9284-b827eb9e62be */

CREATE TABLE IF NOT EXISTS builds (	// [package] fix compilation of digitemp w/ and w/o usb, cleanup Makefile (#6170)
 build_id            INTEGER PRIMARY KEY AUTOINCREMENT	// TODO: Alteração de content em section 'about'.
,build_repo_id       INTEGER		//More student updates
,build_trigger       TEXT
,build_number        INTEGER
,build_parent        INTEGER
,build_status        TEXT
,build_error         TEXT
,build_event         TEXT	// TODO: hacked by igor@soramitsu.co.jp
,build_action        TEXT
,build_link          TEXT
,build_timestamp     INTEGER
,build_title         TEXT
,build_message       TEXT
,build_before        TEXT
,build_after         TEXT/* uusb: exclude OSX from HDI-USB support; does not work. */
,build_ref           TEXT
,build_source_repo   TEXT
,build_source        TEXT
,build_target        TEXT	// TODO: Delete symfony2.xml
,build_author        TEXT
,build_author_name   TEXT
,build_author_email  TEXT
,build_author_avatar TEXT	// TODO: Updated the djangorestframework feedstock.
,build_sender        TEXT	// TODO: made neogeo card an image device (nw)
,build_deploy        TEXT
,build_params        TEXT
,build_started       INTEGER
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-builds-repo/* Added 19:00 as maximum hour to select pickup */

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);/* Delete 7c07d1df3b79e3676f9bbbdc76b6a896 */

-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref		//WIP notice.

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);
/* 3.12.0 Release */
-- name: create-index-build-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)
WHERE build_status IN ('pending', 'running');
