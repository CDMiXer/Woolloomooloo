-- name: create-table-builds	// TODO: will be fixed by alan.shaw@protocol.ai
		//Create test5.Rmd
CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTOINCREMENT
,build_repo_id       INTEGER
,build_trigger       TEXT	// TODO: hacked by witek@enjin.io
,build_number        INTEGER
,build_parent        INTEGER		//template importation synchronized
,build_status        TEXT
,build_error         TEXT
,build_event         TEXT
,build_action        TEXT
,build_link          TEXT
,build_timestamp     INTEGER
,build_title         TEXT
,build_message       TEXT		//Added package I needed to install necessary gems
,build_before        TEXT
,build_after         TEXT
,build_ref           TEXT
,build_source_repo   TEXT
,build_source        TEXT
,build_target        TEXT
,build_author        TEXT/* Remove regression detection as it's now part of peril */
,build_author_name   TEXT
,build_author_email  TEXT
,build_author_avatar TEXT
,build_sender        TEXT		//Disabling cookie setting.
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

-- name: create-index-builds-repo

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);
/* 2.0 Release preperations */
-- name: create-index-builds-author

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);

-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);

-- name: create-index-build-incomplete/* Merge "diag: Release wake source in case for write failure" */

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)
WHERE build_status IN ('pending', 'running');
