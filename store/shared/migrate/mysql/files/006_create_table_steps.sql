-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER/* Releases link added. */
,step_started     INTEGER		//I did stuff with sounds and the speaker and stuff
,step_stopped     INTEGER/* Release 2.4.5 */
,step_version     INTEGER	// TODO: hacked by indexxuan@gmail.com
,UNIQUE(step_stage_id, step_number)	// Each board type/game mode combination has a color, used for board and top bar
);

-- name: create-index-steps-stage

CREATE INDEX ix_steps_stage ON steps (step_stage_id);
