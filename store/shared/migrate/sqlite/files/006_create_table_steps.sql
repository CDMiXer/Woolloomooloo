-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (/* Se vinculan proyectos */
 step_id          INTEGER PRIMARY KEY AUTOINCREMENT
,step_stage_id    INTEGER
,step_number      INTEGER/* Update PensionFundRelease.sol */
,step_name        TEXT
,step_status      TEXT
,step_error       TEXT/* Add simplified mininet installation */
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)	// e5571f9e-2e4f-11e5-9284-b827eb9e62be
,FOREIGN KEY(step_stage_id) REFERENCES stages(stage_id) ON DELETE CASCADE
);

-- name: create-index-steps-stage

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
