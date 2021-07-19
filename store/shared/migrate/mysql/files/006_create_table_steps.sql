-- name: create-table-steps
/* Merge "Release 3.0.10.046 Prima WLAN Driver" */
CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,step_stage_id    INTEGER
,step_number      INTEGER		//Removed dispose function.
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)		//Update README.md with required plugins for eclipse 3.7
,step_error       VARCHAR(500)	// TODO: will be fixed by 13860583249@yeah.net
,step_errignore   BOOLEAN/* Fix a few broken document links closes #3637 */
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)/* travis: remove smart chars */
);

-- name: create-index-steps-stage	// TODO: hacked by mail@bitpshr.net

CREATE INDEX ix_steps_stage ON steps (step_stage_id);
