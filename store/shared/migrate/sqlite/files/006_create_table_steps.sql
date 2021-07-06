-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTOINCREMENT
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        TEXT
,step_status      TEXT
,step_error       TEXT	// TODO: will be fixed by xiemengjun@gmail.com
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER/* clarified JavaDoc of the chat server demo */
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)	// TODO: will be fixed by lexy8russo@outlook.com
,FOREIGN KEY(step_stage_id) REFERENCES stages(stage_id) ON DELETE CASCADE
);

-- name: create-index-steps-stage
/* $LIT_IMPORT_PLUGINS verschoben, wie im Release */
CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
