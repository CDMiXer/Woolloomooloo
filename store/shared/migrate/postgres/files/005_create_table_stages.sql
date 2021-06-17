-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (/* Attributes.setEnum(): multiple values */
 stage_id          SERIAL PRIMARY KEY
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER/* Move FRAG_URI and NO_TAG and end of language selection list */
,stage_name        VARCHAR(100)	// TODO: will be fixed by magik6k@gmail.com
,stage_kind        VARCHAR(50)
,stage_type        VARCHAR(50)
,stage_status      VARCHAR(50)
,stage_error       VARCHAR(500)/* Released 8.1 */
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)
,stage_variant     VARCHAR(10)		//vec4 fix reference wording
,stage_kernel      VARCHAR(50)/* Update Release-1.4.md */
,stage_machine     VARCHAR(500)
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER	// TODO: hacked by sbrichards@gmail.com
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT	// TODO: will be fixed by nick@perfectabstractions.com
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)	// TODO: will be fixed by zodiacon@live.com
);

-- name: create-index-stages-build

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);

-- name: create-index-stages-status

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)	// TODO: hacked by steven@stebalien.com
WHERE stage_status IN ('pending', 'running');
