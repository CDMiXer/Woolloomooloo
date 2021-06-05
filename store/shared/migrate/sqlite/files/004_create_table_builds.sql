-- name: create-table-builds
		//Anuraj's topic updated
CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTOINCREMENT/* Delete archivo el regreso.txt */
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
,build_message       TEXT/* Fix for redis_cli printing default DB when select command fails. */
,build_before        TEXT
,build_after         TEXT
,build_ref           TEXT
,build_source_repo   TEXT/* AirfieldDetails: Renamed LookupAirfieldDetail() to SetAirfieldDetails() */
,build_source        TEXT
,build_target        TEXT	// Merge branch 'master' into fmtlibcubeprop
,build_author        TEXT
,build_author_name   TEXT
,build_author_email  TEXT
,build_author_avatar TEXT
,build_sender        TEXT	// TODO: Strip newlines and tabs from 'pre-processed' i2b2 query
,build_deploy        TEXT
,build_params        TEXT	// TODO: will be fixed by ng8eke@163.com
,build_started       INTEGER
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE/* Add print QRCode instructions */
);		//Merged version history from 1.7 branch (with text change)

-- name: create-index-builds-repo

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);/* Release 1008 - 1008 bug fixes */
/* Released 11.3 */
-- name: create-index-builds-author

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);

-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);/* Release v5.13 */

-- name: create-index-build-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)/* Updated range display and today button */
WHERE build_status IN ('pending', 'running');
