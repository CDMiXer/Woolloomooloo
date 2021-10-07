-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (/* Release version 0.9.0 */
 step_id          SERIAL PRIMARY KEY
,step_stage_id    INTEGER		//fix for start/end dates default values to be 1 month in range
,step_number      INTEGER/* update to 2.27.x Release Candidate 2 (2.27.2) */
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER/* Release of eeacms/eprtr-frontend:0.3-beta.25 */
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);

-- name: create-index-steps-stage/* Add jemoji dependecies */

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
