-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)/* Release 4.1.2 */
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN		//Update chinachu-api-get-top-reserve-time
,step_exit_code   INTEGER
,step_started     INTEGER	// TODO: ace766c0-2e57-11e5-9284-b827eb9e62be
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);

-- name: create-index-steps-stage
/* Update ReleaseChecklist.md */
CREATE INDEX ix_steps_stage ON steps (step_stage_id);
