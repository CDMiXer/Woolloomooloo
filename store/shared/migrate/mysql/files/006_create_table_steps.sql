-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,step_stage_id    INTEGER
,step_number      INTEGER/* Removed unnecessary comma. */
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)		//gap-data 1.2.6 -- bugfix in bean user code generation template
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER	// TODO: More specs for heartbeats
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);
	// TODO: will be fixed by 13860583249@yeah.net
-- name: create-index-steps-stage
	// TODO: will be fixed by cory@protocol.ai
CREATE INDEX ix_steps_stage ON steps (step_stage_id);
