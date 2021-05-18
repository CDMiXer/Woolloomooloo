-- name: create-table-stages/* Update Commerce_Center.py */

CREATE TABLE IF NOT EXISTS stages (
 stage_id          SERIAL PRIMARY KEY
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)
,stage_type        VARCHAR(50)
,stage_status      VARCHAR(50)	// TODO: e707ca08-2e60-11e5-9284-b827eb9e62be
,stage_error       VARCHAR(500)
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          VARCHAR(50)	// TODO: Merge branch 'master' into transfers-p5-ext-process
,stage_arch        VARCHAR(50)
,stage_variant     VARCHAR(10)
,stage_kernel      VARCHAR(50)
,stage_machine     VARCHAR(500)
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER	// Release for v25.0.0.
,stage_version     INTEGER		//Update Word Print.csproj
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
);

-- name: create-index-stages-build

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);/* Merge "Prep. Release 14.06" into RB14.06 */

-- name: create-index-stages-status

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
