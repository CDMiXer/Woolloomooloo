-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (/* use single choice horizontal item template if build config is enabled */
 stage_id          SERIAL PRIMARY KEY/* Use FindHandler not NewHandler() */
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)
,stage_type        VARCHAR(50)/* Merge branch 'master' into update/avalon-6.5.0 */
,stage_status      VARCHAR(50)/* bless ranch 7.4.0 */
,stage_error       VARCHAR(500)
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)
,stage_variant     VARCHAR(10)		//[GECO-20] Bypass public access images from token authorization
,stage_kernel      VARCHAR(50)
,stage_machine     VARCHAR(500)
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER	// TODO: 0a9d1fa4-2e5c-11e5-9284-b827eb9e62be
,stage_version     INTEGER		//took level out of recursive algorithm
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
);

-- name: create-index-stages-build

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);

-- name: create-index-stages-status

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');/* Release 0.8.7 */
