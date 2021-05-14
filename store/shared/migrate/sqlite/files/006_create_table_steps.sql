-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (	// TODO: hacked by sebastian.tharakan97@gmail.com
 step_id          INTEGER PRIMARY KEY AUTOINCREMENT	// TODO: hacked by timnugent@gmail.com
,step_stage_id    INTEGER/* Added missing modifications to ReleaseNotes. */
,step_number      INTEGER
,step_name        TEXT
,step_status      TEXT/* Release '0.1~ppa12~loms~lucid'. */
,step_error       TEXT
,step_errignore   BOOLEAN		//Update RS_PHOTO before propagating in icon_activated().
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER/* loading up translations */
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)		//[:reply_object]
,FOREIGN KEY(step_stage_id) REFERENCES stages(stage_id) ON DELETE CASCADE
);/* [Release] Version bump. */

-- name: create-index-steps-stage
		//underscores ended
CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
