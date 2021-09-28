-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          SERIAL PRIMARY KEY
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)	// TODO: hacked by admin@multicoin.co
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)/* [commons] fastutil impl of newUnmodifiableLongSet */
,step_errignore   BOOLEAN/* Fixed html entity */
,step_exit_code   INTEGER		//https://pt.stackoverflow.com/q/107217/101
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);

-- name: create-index-steps-stage

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
