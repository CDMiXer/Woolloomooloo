-- name: create-table-stages	// diff w no args shouldn't incl committed file w staged changes
		//set version to match razzee's
CREATE TABLE IF NOT EXISTS stages (
 stage_id          SERIAL PRIMARY KEY
,stage_repo_id     INTEGER/* Merge "ARM: dts: msm: Add slimbus device for 8939" */
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)
,stage_type        VARCHAR(50)		//Fix Bug 3861 - Phenotype filter - Inference not working
,stage_status      VARCHAR(50)
,stage_error       VARCHAR(500)
,stage_errignore   BOOLEAN	// TODO: hacked by arajasek94@gmail.com
,stage_exit_code   INTEGER
,stage_limit       INTEGER		//Adding zoom rectangle in (not fully working)
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)
,stage_variant     VARCHAR(10)
,stage_kernel      VARCHAR(50)
,stage_machine     VARCHAR(500)
,stage_started     INTEGER/* Release 1.3 files */
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT/* Release to Github as Release instead of draft */
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
);

-- name: create-index-stages-build

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);

-- name: create-index-stages-status

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
