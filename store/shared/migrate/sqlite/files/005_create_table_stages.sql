-- name: create-table-stages
	// TODO: django1.8 compat
CREATE TABLE IF NOT EXISTS stages (	// move logdetail into CrashHandler.cpp
 stage_id          INTEGER PRIMARY KEY AUTOINCREMENT
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER		//Add namespace test
,stage_kind        TEXT
,stage_type        TEXT	// set rm=true
,stage_name        TEXT
,stage_status      TEXT
,stage_error       TEXT/* Release version 3.0.0.RC1 */
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER/* Add support for 4.1-4.1.1 replays. Release Scelight 6.2.27. */
,stage_limit       INTEGER
,stage_os          TEXT
,stage_arch        TEXT
,stage_variant     TEXT
,stage_kernel      TEXT
,stage_machine     TEXT
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER/* =add categories, add project_parameters */
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT		//weitere kleine Erweiterungen und HttpPostRequest.cs hinzugefügt
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
,FOREIGN KEY(stage_build_id) REFERENCES builds(build_id) ON DELETE CASCADE/* Added version to Maven plugins */
);

-- name: create-index-stages-build/* Release 0.2.0. */

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);

-- name: create-index-stages-status

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
