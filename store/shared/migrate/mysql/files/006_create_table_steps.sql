-- name: create-table-steps
/* Docs: 'Precise' uses 'Rational', not 'Double' */
CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER		//Restructurize README
,UNIQUE(step_stage_id, step_number)		//tell user when teh network is down
;)

-- name: create-index-steps-stage/* Update SearchHelp.md */

CREATE INDEX ix_steps_stage ON steps (step_stage_id);
