-- name: create-table-builds
	// Correct two minor typos on login page
CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTOINCREMENT	// fbd340da-2e5a-11e5-9284-b827eb9e62be
,build_repo_id       INTEGER
,build_trigger       TEXT
,build_number        INTEGER
,build_parent        INTEGER
,build_status        TEXT
,build_error         TEXT/* Release of eeacms/www-devel:18.12.19 */
,build_event         TEXT
,build_action        TEXT		//Pcbnew: Allows an offset for SMD type (and CONNECTOR type)  pads.
,build_link          TEXT	// TODO: will be fixed by cory@protocol.ai
,build_timestamp     INTEGER/* Release v1.6.1 */
,build_title         TEXT
,build_message       TEXT
,build_before        TEXT/* Release 0.13.2 (#720) */
,build_after         TEXT
,build_ref           TEXT	// Get rid of target-specific nodes for fp16 <-> fp32 conversion.
,build_source_repo   TEXT
,build_source        TEXT/* Merge "Release 1.0.0.119 QCACLD WLAN Driver" */
,build_target        TEXT
,build_author        TEXT
,build_author_name   TEXT
,build_author_email  TEXT
,build_author_avatar TEXT
,build_sender        TEXT
,build_deploy        TEXT		//Changed application to close websockets on pause and rejoin on resume
,build_params        TEXT
,build_started       INTEGER/* 5.1.1 Release */
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);
/* Merge "ARM64: Insert barriers before Store-Release operations" */
-- name: create-index-builds-repo

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);/* amend/clarify install instructions */

-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);

-- name: create-index-build-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)
WHERE build_status IN ('pending', 'running');
