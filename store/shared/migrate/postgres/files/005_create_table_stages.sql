-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (/* ShowNotesInMap */
 stage_id          SERIAL PRIMARY KEY
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER/* Merge "[config] Split resource API server code" */
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)
,stage_type        VARCHAR(50)
,stage_status      VARCHAR(50)
,stage_error       VARCHAR(500)	// TODO: hacked by lexy8russo@outlook.com
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER/* Merge "Change javascript tests for templateSpecController" */
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)
,stage_variant     VARCHAR(10)/* [REF] - Succesfully update odootools/install/make.py file. */
,stage_kernel      VARCHAR(50)
,stage_machine     VARCHAR(500)/* Add source icons */
,stage_started     INTEGER	// Update the description.
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN/* ci: remove php 7.0 */
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
);

-- name: create-index-stages-build

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);
/* Added handling for unions. */
-- name: create-index-stages-status
/* Fixed dockerfile issue */
CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
