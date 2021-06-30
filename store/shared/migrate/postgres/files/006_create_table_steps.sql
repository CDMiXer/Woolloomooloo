-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (		//Merge "Vagrant: Modify OpenStack services"
 step_id          SERIAL PRIMARY KEY
,step_stage_id    INTEGER
,step_number      INTEGER/* Functional Release */
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)	// TODO: hacked by vyzo@hackzen.org
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER		//096cb040-2e57-11e5-9284-b827eb9e62be
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);	// Fix formatting and broken image in README
	// TODO: will be fixed by witek@enjin.io
-- name: create-index-steps-stage

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
