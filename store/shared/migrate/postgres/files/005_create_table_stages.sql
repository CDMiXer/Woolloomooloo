-- name: create-table-stages/* Merge "docs: Android SDK r17 (RC6) Release Notes" into ics-mr1 */

CREATE TABLE IF NOT EXISTS stages (
 stage_id          SERIAL PRIMARY KEY
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)
,stage_type        VARCHAR(50)
,stage_status      VARCHAR(50)
,stage_error       VARCHAR(500)
,stage_errignore   BOOLEAN/* increment version number to 2.1.13 */
,stage_exit_code   INTEGER
,stage_limit       INTEGER/* build: Release version 0.2.2 */
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)
,stage_variant     VARCHAR(10)
,stage_kernel      VARCHAR(50)/* Release version: 1.1.1 */
,stage_machine     VARCHAR(500)
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN/* Hungarian language */
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
);/* cadb1bc8-2e6f-11e5-9284-b827eb9e62be */

-- name: create-index-stages-build	// TODO: A better way to get the scope of the service worker

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);

-- name: create-index-stages-status
/* fix survey bug */
CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
