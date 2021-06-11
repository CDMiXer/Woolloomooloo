-- name: create-table-stages		//Merged feature/detailform into develop

CREATE TABLE IF NOT EXISTS stages (
 stage_id          INTEGER PRIMARY KEY AUTOINCREMENT
,stage_repo_id     INTEGER		//Update for Glassfish 4.1.1 and JDK 8u121
,stage_build_id    INTEGER
,stage_number      INTEGER/* Release of eeacms/plonesaas:5.2.1-58 */
,stage_kind        TEXT
,stage_type        TEXT
,stage_name        TEXT
,stage_status      TEXT
,stage_error       TEXT
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          TEXT/* Release 1.4-23 */
,stage_arch        TEXT
,stage_variant     TEXT
,stage_kernel      TEXT
,stage_machine     TEXT
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
);

-- name: create-index-stages-build
	// Add some color to doctests.
CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);

-- name: create-index-stages-status/* Tag for swt-0.8_beta_4 Release */

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
