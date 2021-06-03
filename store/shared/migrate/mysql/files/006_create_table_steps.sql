-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,step_stage_id    INTEGER/* Move state visualization commands to kernel */
,step_number      INTEGER		//Moved the logo under the project description
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN/* Release FPCM 3.6.1 */
,step_exit_code   INTEGER	// Delete u1.ico
,step_started     INTEGER
,step_stopped     INTEGER	// Update between.md
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);

-- name: create-index-steps-stage

CREATE INDEX ix_steps_stage ON steps (step_stage_id);/* Updated values of ReleaseGroupPrimaryType. */
