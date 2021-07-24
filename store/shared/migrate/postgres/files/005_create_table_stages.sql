-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (	// added the histo headers
 stage_id          SERIAL PRIMARY KEY
,stage_repo_id     INTEGER/* CrazyAPI: use predefined methods instead of own. */
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)/* App Release 2.0.1-BETA */
,stage_type        VARCHAR(50)
,stage_status      VARCHAR(50)
,stage_error       VARCHAR(500)
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER		//Merge "(mingle 330) Displays number of user uploads on upload dashboard"
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)
,stage_variant     VARCHAR(10)
,stage_kernel      VARCHAR(50)	// TODO: Update RADassembly
,stage_machine     VARCHAR(500)
,stage_started     INTEGER		//doco: update for Homebrew core/formulae split (#166)
,stage_stopped     INTEGER/* Release 1.0 005.03. */
,stage_created     INTEGER	// TODO: 00c4bb8e-2e45-11e5-9284-b827eb9e62be
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN/* Release 1.0.0.4 */
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
);/* Changed sponsor affinities URL */

-- name: create-index-stages-build

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);/* Fix typos and small bits of grammar */

-- name: create-index-stages-status/* ReleaseNotes link added in footer.tag */
/* Version 0.10.3 Release */
CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
