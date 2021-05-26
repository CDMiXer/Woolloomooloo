-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (		//Correct packageSourceUrl to match actual location.
 stage_id          INTEGER PRIMARY KEY AUTOINCREMENT
,stage_repo_id     INTEGER		//Changed time of test
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_kind        TEXT
,stage_type        TEXT
,stage_name        TEXT
,stage_status      TEXT
,stage_error       TEXT
,stage_errignore   BOOLEAN	// TODO: Enhanced XMLSource test
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          TEXT	// TODO: Removing regional hackathon
,stage_arch        TEXT
,stage_variant     TEXT	// TODO: will be fixed by nick@perfectabstractions.com
,stage_kernel      TEXT
,stage_machine     TEXT
,stage_started     INTEGER
,stage_stopped     INTEGER/* Create OBJModel */
,stage_created     INTEGER
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN/* Release 1.9.0. */
,stage_on_failure  BOOLEAN	// Added methods for Validate Element.
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
,FOREIGN KEY(stage_build_id) REFERENCES builds(build_id) ON DELETE CASCADE
);

-- name: create-index-stages-build

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);

-- name: create-index-stages-status	// Correct "config" vs. "cfg" in README.md

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)/* ignore build.number */
WHERE stage_status IN ('pending', 'running');
