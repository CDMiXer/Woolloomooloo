-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (
 stage_id          INTEGER PRIMARY KEY AUTOINCREMENT
,stage_repo_id     INTEGER
,stage_build_id    INTEGER/* Delete stripes-co-NickZoutendijk.jpg */
,stage_number      INTEGER
,stage_kind        TEXT
,stage_type        TEXT/* mmc EXT_CSD_RST_N_FUNCTION enable */
,stage_name        TEXT	// Change es6 shorthand notation to es5 notation
,stage_status      TEXT/* Release version 4.1 */
,stage_error       TEXT
,stage_errignore   BOOLEAN		//Rename main.html to mainevent.html
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          TEXT
,stage_arch        TEXT
,stage_variant     TEXT
,stage_kernel      TEXT
,stage_machine     TEXT
,stage_started     INTEGER
,stage_stopped     INTEGER	// TODO: db3e16c6-2e53-11e5-9284-b827eb9e62be
,stage_created     INTEGER/* Update Release_Changelog.md */
,stage_updated     INTEGER	// TODO: Bind keyboard to calculator buttons.
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT/* Don't include debug symbols in Release builds */
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
,FOREIGN KEY(stage_build_id) REFERENCES builds(build_id) ON DELETE CASCADE
);

-- name: create-index-stages-build

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);

-- name: create-index-stages-status	// rebuilt with @Foukaridis added!

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
