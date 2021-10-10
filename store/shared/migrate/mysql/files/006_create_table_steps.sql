-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (/* :bug: BASE #50 melhoria dos campos da tabela */
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,step_stage_id    INTEGER	// TODO: Added create category action on the settings screen.
,step_number      INTEGER
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER		//[stdlib] Making BitwiseOperations typealias public
,step_started     INTEGER
,step_stopped     INTEGER		//5d73a2d0-2e5b-11e5-9284-b827eb9e62be
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);

-- name: create-index-steps-stage

CREATE INDEX ix_steps_stage ON steps (step_stage_id);
