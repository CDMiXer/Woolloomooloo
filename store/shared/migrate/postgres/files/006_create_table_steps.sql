-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          SERIAL PRIMARY KEY
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)	// TODO: Merge "Grant HeifWriterTest read/write permission by rules" into pi-androidx-dev
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER		//Implementing additional methods to attach data to output file.
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);/* CyFluxViz Release v0.88. */
	// TODO: hacked by boringland@protonmail.ch
-- name: create-index-steps-stage

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
