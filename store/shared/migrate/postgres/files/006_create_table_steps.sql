-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          SERIAL PRIMARY KEY
,step_stage_id    INTEGER		//[README] Added circleci badge
,step_number      INTEGER
,step_name        VARCHAR(100)/* Release version 0.4.1 */
,step_status      VARCHAR(50)		//Update ru.inf
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER/* Release of eeacms/eprtr-frontend:0.2-beta.23 */
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);	// TODO: hacked by hello@brooklynzelenka.com

-- name: create-index-steps-stage

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
