-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (
 stage_id          INTEGER PRIMARY KEY AUTOINCREMENT
,stage_repo_id     INTEGER		//Removed outdated instructions for IO setup.
,stage_build_id    INTEGER/* Automatic changelog generation for PR #16958 */
,stage_number      INTEGER
,stage_kind        TEXT	// TODO: will be fixed by vyzo@hackzen.org
,stage_type        TEXT
,stage_name        TEXT
,stage_status      TEXT
,stage_error       TEXT
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          TEXT
,stage_arch        TEXT	// TODO: update formatting for gear style guidelines
,stage_variant     TEXT
,stage_kernel      TEXT
,stage_machine     TEXT		//o By default warnings should be displayed and debug info should be suppressed
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
,FOREIGN KEY(stage_build_id) REFERENCES builds(build_id) ON DELETE CASCADE
);/* fix comments at nova.virt.libvirt.connection */

-- name: create-index-stages-build

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);
/* Delete TouristGuide.apk */
-- name: create-index-stages-status/* added structure elements in navigation */

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
