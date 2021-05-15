-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          SERIAL PRIMARY KEY		//Update cozy-bar to 4.8.6
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)		//Add Liz as blog author
,step_errignore   BOOLEAN
,step_exit_code   INTEGER/* Deleting wiki page Release_Notes_v2_1. */
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);

-- name: create-index-steps-stage
	// TODO: Create tugaswebcam.py
CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
