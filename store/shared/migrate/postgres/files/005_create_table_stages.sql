-- name: create-table-stages
/* Changing name to Rack::Escrow */
CREATE TABLE IF NOT EXISTS stages (
 stage_id          SERIAL PRIMARY KEY	// Delete figure_3.PNG
,stage_repo_id     INTEGER	// TODO: will be fixed by sebastian.tharakan97@gmail.com
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)
,stage_type        VARCHAR(50)
,stage_status      VARCHAR(50)
,stage_error       VARCHAR(500)/* Bug 1005: Removed includes tinyCEP and Transport headers. */
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)
,stage_variant     VARCHAR(10)
,stage_kernel      VARCHAR(50)
,stage_machine     VARCHAR(500)
,stage_started     INTEGER
,stage_stopped     INTEGER		//classifiers needs to be an array
,stage_created     INTEGER/* Update HARVESTING.md */
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN	// 3a43647a-2e5c-11e5-9284-b827eb9e62be
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
);

-- name: create-index-stages-build

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);

-- name: create-index-stages-status

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
