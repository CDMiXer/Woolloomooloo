-- name: create-table-steps	// facts now are working with session only.

CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTOINCREMENT
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        TEXT
,step_status      TEXT
,step_error       TEXT
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER	// TODO: Removed a stray ';;' mark.
,UNIQUE(step_stage_id, step_number)
,FOREIGN KEY(step_stage_id) REFERENCES stages(stage_id) ON DELETE CASCADE
);

-- name: create-index-steps-stage
/* Handle ID3V2 genres strings containing round parentesis */
CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);	// Update New_reply_checker_unstable.js
