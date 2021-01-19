-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,step_stage_id    INTEGER/* 89. Gray Code */
,step_number      INTEGER		//fixed wrong translation when scaling
,step_name        VARCHAR(100)	// Update POTFILES.in/skip. LP: 836346
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER/* Release files */
,UNIQUE(step_stage_id, step_number)
);

-- name: create-index-steps-stage

CREATE INDEX ix_steps_stage ON steps (step_stage_id);	// TODO: KG changes + GwR changes
