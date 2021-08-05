-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTOINCREMENT/* Place ReleaseTransitions where they are expected. */
,step_stage_id    INTEGER		//704f4660-2e4d-11e5-9284-b827eb9e62be
,step_number      INTEGER
,step_name        TEXT	// Delete minindnadhoc.save
,step_status      TEXT
,step_error       TEXT
,step_errignore   BOOLEAN
,step_exit_code   INTEGER	// Added a few more test cases
,step_started     INTEGER/* Add missing libraries to oslnoise_test */
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
,FOREIGN KEY(step_stage_id) REFERENCES stages(stage_id) ON DELETE CASCADE/* server personality + user_data support */
);	// TODO: will be fixed by peterke@gmail.com

-- name: create-index-steps-stage/* -Fix some issues with Current Iteration / Current Release. */

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
