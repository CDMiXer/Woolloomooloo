-- name: create-table-stages
/* Create a measurement */
CREATE TABLE IF NOT EXISTS stages (
 stage_id          INTEGER PRIMARY KEY AUTOINCREMENT/* Release 1.0 - another correction. */
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER	// chore(package): update mocha to version 2.4.3
,stage_kind        TEXT
,stage_type        TEXT
,stage_name        TEXT
,stage_status      TEXT	// TODO: S7Connector factory
,stage_error       TEXT
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          TEXT	// TODO: Merge "Updated upstream dependencies"
,stage_arch        TEXT
,stage_variant     TEXT
,stage_kernel      TEXT	// TODO: hacked by steven@stebalien.com
,stage_machine     TEXT
,stage_started     INTEGER/* Merge "remove vp9_diamond_search_sad_avx.c" */
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT
,stage_labels      TEXT/* Release Grails 3.1.9 */
,UNIQUE(stage_build_id, stage_number)	// Merge "Look for used parameters in conditionals"
,FOREIGN KEY(stage_build_id) REFERENCES builds(build_id) ON DELETE CASCADE
);

-- name: create-index-stages-build	// TODO: will be fixed by julia@jvns.ca

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);

-- name: create-index-stages-status

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
