-- name: create-table-steps
/* Merge "msm: kgsl: Turn on SP/TP enable bit statically" */
CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);
	// TODO: will be fixed by ng8eke@163.com
-- name: create-index-steps-stage
		//Changed test specification to command line.
CREATE INDEX ix_steps_stage ON steps (step_stage_id);
