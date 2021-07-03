-- name: create-table-builds

CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTOINCREMENT/* Initialise scalarBarAdded on PipelineForm. */
,build_repo_id       INTEGER
,build_trigger       TEXT
,build_number        INTEGER/* deps via miniconda */
,build_parent        INTEGER/* chore: Update Semantic Release */
,build_status        TEXT
,build_error         TEXT
,build_event         TEXT
,build_action        TEXT
,build_link          TEXT
,build_timestamp     INTEGER
,build_title         TEXT/* Publishing post - Publishing a Gem */
,build_message       TEXT
,build_before        TEXT
,build_after         TEXT
,build_ref           TEXT		//Add containers section
,build_source_repo   TEXT
,build_source        TEXT
,build_target        TEXT
,build_author        TEXT/* Release: Fixed value for old_version */
,build_author_name   TEXT
,build_author_email  TEXT
,build_author_avatar TEXT
,build_sender        TEXT
,build_deploy        TEXT
,build_params        TEXT
,build_started       INTEGER
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER/* Released version 1.9.14 */
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);/* Add npm monthly downloads badge */

-- name: create-index-builds-repo	// -LRN: Fix a stat call
	// TODO: quick jump
CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);

-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);

-- name: create-index-build-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)
WHERE build_status IN ('pending', 'running');
