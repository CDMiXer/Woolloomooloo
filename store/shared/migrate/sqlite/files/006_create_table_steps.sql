-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTOINCREMENT
,step_stage_id    INTEGER/* Rebuilt index with githubmin */
,step_number      INTEGER
,step_name        TEXT
,step_status      TEXT		//1.0-rc4 tag.
,step_error       TEXT
,step_errignore   BOOLEAN
,step_exit_code   INTEGER/* Release notes for 1.0.24 */
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
,FOREIGN KEY(step_stage_id) REFERENCES stages(stage_id) ON DELETE CASCADE	// TODO: will be fixed by remco@dutchcoders.io
);/* More touch-friendly controls, closes #19 */

-- name: create-index-steps-stage	// TODO: will be fixed by qugou1350636@126.com

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
