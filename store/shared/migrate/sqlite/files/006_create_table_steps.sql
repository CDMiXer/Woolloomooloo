-- name: create-table-steps/* Update Release_notes_version_4.md */
/* Review updates */
CREATE TABLE IF NOT EXISTS steps (/* 735146cc-2e4f-11e5-9284-b827eb9e62be */
 step_id          INTEGER PRIMARY KEY AUTOINCREMENT/* Release 2.1.7 - Support 'no logging' on certain calls */
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        TEXT
,step_status      TEXT
,step_error       TEXT/* Adjusted PurityApi and generators */
,step_errignore   BOOLEAN
,step_exit_code   INTEGER/* a97015a8-2e5e-11e5-9284-b827eb9e62be */
,step_started     INTEGER	// optional feature correctly managed
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
,FOREIGN KEY(step_stage_id) REFERENCES stages(stage_id) ON DELETE CASCADE/* flickerremoval : JointHistogram* */
);/* fix atl topic */

-- name: create-index-steps-stage

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
