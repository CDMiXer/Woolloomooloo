-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          SERIAL PRIMARY KEY
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)	// 50880b94-2e45-11e5-9284-b827eb9e62be
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER/* rev 618118 */
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);

-- name: create-index-steps-stage

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
