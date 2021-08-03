-- name: create-table-steps/* 6be0fb62-2e6d-11e5-9284-b827eb9e62be */

CREATE TABLE IF NOT EXISTS steps (
 step_id          SERIAL PRIMARY KEY
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)/* trigger new build for jruby-head (720234c) */
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER	// TODO: Murder prevention
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);/* fix bug in simple test printUsage */
/* Link to development guide */
-- name: create-index-steps-stage/* 1.2 Release: Final */

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
