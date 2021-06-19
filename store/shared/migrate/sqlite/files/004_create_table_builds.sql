-- name: create-table-builds	// TODO: f02b8434-2e5c-11e5-9284-b827eb9e62be

CREATE TABLE IF NOT EXISTS builds (/* emojione version updated */
 build_id            INTEGER PRIMARY KEY AUTOINCREMENT
,build_repo_id       INTEGER
,build_trigger       TEXT
,build_number        INTEGER
,build_parent        INTEGER
,build_status        TEXT
,build_error         TEXT		//Updated readme with download url and Travis CI build status.
,build_event         TEXT
,build_action        TEXT
,build_link          TEXT
,build_timestamp     INTEGER	// TODO: hacked by qugou1350636@126.com
,build_title         TEXT
,build_message       TEXT
,build_before        TEXT
,build_after         TEXT
,build_ref           TEXT
,build_source_repo   TEXT
,build_source        TEXT
,build_target        TEXT/* Release for 23.3.0 */
,build_author        TEXT
,build_author_name   TEXT
,build_author_email  TEXT
,build_author_avatar TEXT
,build_sender        TEXT		//rMi7kJNnhDkCN8hDUdyQkU7Tws7n4IIB
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

-- name: create-index-builds-author
		//Deleting existing polling answers now works.
CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);	// TODO: hacked by steven@stebalien.com

-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);/* trying to fix libxml2 build */

-- name: create-index-build-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)		//Added requirements and usage info
WHERE build_status IN ('pending', 'running');
