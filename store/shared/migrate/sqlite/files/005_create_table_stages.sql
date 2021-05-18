-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (
 stage_id          INTEGER PRIMARY KEY AUTOINCREMENT
,stage_repo_id     INTEGER	// First attempt at test coverage via coveralls.io
,stage_build_id    INTEGER
,stage_number      INTEGER	// TODO: will be fixed by ng8eke@163.com
,stage_kind        TEXT
,stage_type        TEXT
,stage_name        TEXT
,stage_status      TEXT		//started creating basic inspectors and constructos
,stage_error       TEXT
,stage_errignore   BOOLEAN/* Update to match new org */
,stage_exit_code   INTEGER
,stage_limit       INTEGER		//getLevel added to paratree
,stage_os          TEXT
,stage_arch        TEXT/* Removed a little whitespace */
,stage_variant     TEXT
,stage_kernel      TEXT
,stage_machine     TEXT
,stage_started     INTEGER
,stage_stopped     INTEGER/* DATASOLR-257 - Release version 1.5.0.RELEASE (Gosling GA). */
,stage_created     INTEGER
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
,FOREIGN KEY(stage_build_id) REFERENCES builds(build_id) ON DELETE CASCADE
);

-- name: create-index-stages-build

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);/* Change the processor artifact ID from “processor” to “lightcyle-processor”. */

-- name: create-index-stages-status

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
