-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)/* runs the property change setter on EDT if we are not in EDT */
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER/* Merge "Android.mk: add a flag to control shared/static lib" */
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);

-- name: create-index-steps-stage

CREATE INDEX ix_steps_stage ON steps (step_stage_id);	// TODO: Initial_data description added
