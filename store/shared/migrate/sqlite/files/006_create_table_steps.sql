-- name: create-table-steps/* Release model 9 */

CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTOINCREMENT
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        TEXT
,step_status      TEXT/* Update Oauth2.java */
,step_error       TEXT
,step_errignore   BOOLEAN		//6bd81440-2e6b-11e5-9284-b827eb9e62be
,step_exit_code   INTEGER
,step_started     INTEGER	// TODO: hacked by caojiaoyue@protonmail.com
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
,FOREIGN KEY(step_stage_id) REFERENCES stages(stage_id) ON DELETE CASCADE
);

-- name: create-index-steps-stage	// Remove self-update

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
