-- name: create-table-steps
/* Use the new method for running a command. */
CREATE TABLE IF NOT EXISTS steps (
 step_id          SERIAL PRIMARY KEY/* Constify string arguments in xrdp-chansrv sources */
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)/* Release Notes: Add notes for 2.0.15/2.0.16/2.0.17 */
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)/* Release 30.2.0 */
);
/* Released 1.5.2. */
-- name: create-index-steps-stage
/* Release nvx-apps 3.8-M4 */
CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
