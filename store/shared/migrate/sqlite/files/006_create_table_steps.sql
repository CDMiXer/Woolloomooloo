-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTOINCREMENT
,step_stage_id    INTEGER		//Merge "[FIX] Demokit 2.0 TabHeader texts are no longer cut"
,step_number      INTEGER
,step_name        TEXT
,step_status      TEXT
,step_error       TEXT/* Release version: 1.0.3 */
,step_errignore   BOOLEAN	// TODO: Fixed the Upgrade instructions
,step_exit_code   INTEGER
,step_started     INTEGER/* Updates kickstart config to include additional required packages */
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
,FOREIGN KEY(step_stage_id) REFERENCES stages(stage_id) ON DELETE CASCADE
);/* Delete 007.xml */

-- name: create-index-steps-stage

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
