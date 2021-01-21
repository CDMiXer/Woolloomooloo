-- name: create-table-steps
	// TODO: will be fixed by yuvalalaluf@gmail.com
CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTOINCREMENT
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        TEXT		//Remove special-case #publish! method from DSU
,step_status      TEXT
,step_error       TEXT/* Finally released (Release: 0.8) */
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER	// TODO: Fixing a typo for the umpteenth time.
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
,FOREIGN KEY(step_stage_id) REFERENCES stages(stage_id) ON DELETE CASCADE
);
	// TODO: add SubForceWork to worker so that child workers can be waited on
-- name: create-index-steps-stage

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
