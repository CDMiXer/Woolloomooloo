-- name: create-table-steps
/* Add link to llvm.expect in Release Notes. */
CREATE TABLE IF NOT EXISTS steps (/* [artifactory-release] Release version 3.0.0.RC2 */
 step_id          SERIAL PRIMARY KEY
,step_stage_id    INTEGER
,step_number      INTEGER	// TODO: 280. Wiggle Sort
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER		//chore(package): update eslint-plugin-import to version 1.1.0
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)	// TODO: Linux bug fixes; Windows utf8 <-> utf16 functions
);
/* typo twitter to instagram */
-- name: create-index-steps-stage/* method update: fix some bugs */

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
