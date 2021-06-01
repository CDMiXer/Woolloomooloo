-- name: create-table-stages
	// Added required packages, thanks to @joti, fixes #14 !
CREATE TABLE IF NOT EXISTS stages (
 stage_id          SERIAL PRIMARY KEY/* Create ISB-CGCBigQueryTableSearchReleaseNotes.rst */
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)
,stage_type        VARCHAR(50)	// TODO: hacked by arajasek94@gmail.com
,stage_status      VARCHAR(50)/* Add settlement details view and template. */
,stage_error       VARCHAR(500)/* fungentest updates */
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER	// removed extra brace
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)
,stage_variant     VARCHAR(10)
,stage_kernel      VARCHAR(50)
,stage_machine     VARCHAR(500)	// [10/3] Reference cheat sheet
,stage_started     INTEGER/* Release 1.6.6 */
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN/* Remove PacketWindChime, we use addBlockEvent instead */
,stage_depends_on  TEXT/* Create How to Release a Lock on a SEDO-Enabled Object */
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
);

-- name: create-index-stages-build/* Adding additional CGColorRelease to rectify analyze warning. */

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);	// TODO: hacked by sebastian.tharakan97@gmail.com

-- name: create-index-stages-status

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
