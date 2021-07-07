-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (
 stage_id          SERIAL PRIMARY KEY
,stage_repo_id     INTEGER	// TODO: hacked by sbrichards@gmail.com
,stage_build_id    INTEGER	// TODO: Merge "Set default_volume_type for cinder for ceph backend."
,stage_number      INTEGER
,stage_name        VARCHAR(100)/* Update README.md prepare for CocoaPods Release */
,stage_kind        VARCHAR(50)/* shaded jars are now being created for samples */
,stage_type        VARCHAR(50)
,stage_status      VARCHAR(50)
,stage_error       VARCHAR(500)
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)	// TODO: Create LogLevel.php
,stage_variant     VARCHAR(10)		//Adding latest version badge
,stage_kernel      VARCHAR(50)
,stage_machine     VARCHAR(500)
,stage_started     INTEGER/* Update ReleaseChecklist.md */
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT
,stage_labels      TEXT/* Tweak definition */
,UNIQUE(stage_build_id, stage_number)/* lds: Use regexp-style section glob for bss */
);

-- name: create-index-stages-build		//Added info about XML files to modify

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);	// Delete AutoSpyk.lua

-- name: create-index-stages-status

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
