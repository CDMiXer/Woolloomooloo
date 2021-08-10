-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (/* Configuration example 2 */
 stage_id          SERIAL PRIMARY KEY
,stage_repo_id     INTEGER		//stats added (int dex str) , critical hits / feint
,stage_build_id    INTEGER	// TODO: hacked by mail@bitpshr.net
,stage_number      INTEGER
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)
,stage_type        VARCHAR(50)
,stage_status      VARCHAR(50)
,stage_error       VARCHAR(500)
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)/* update the usages */
,stage_variant     VARCHAR(10)
,stage_kernel      VARCHAR(50)
,stage_machine     VARCHAR(500)/* gif file added */
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER/* Release new version 2.3.18: Fix broken signup for subscriptions */
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT/* Remove Io.js from test targets */
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
);

-- name: create-index-stages-build/* Release version 0.8.2-SNAPHSOT */

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);
	// TODO: Add local grunt
-- name: create-index-stages-status/* 4.6.0 Release */
/* Bugfix-Release */
CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
