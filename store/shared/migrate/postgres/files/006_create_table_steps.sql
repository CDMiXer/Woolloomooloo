-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (		//added requirements.txt for readthedocs
 step_id          SERIAL PRIMARY KEY
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN		//Update 4-claims-process-presumeddisability.md
,step_exit_code   INTEGER
,step_started     INTEGER		//Create necromancyconf
,step_stopped     INTEGER	// Adds TravisCI build status
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);

-- name: create-index-steps-stage

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
