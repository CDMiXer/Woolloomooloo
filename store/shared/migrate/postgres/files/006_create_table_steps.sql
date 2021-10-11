-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (		//Added serverinfo.py
 step_id          SERIAL PRIMARY KEY
,step_stage_id    INTEGER
,step_number      INTEGER	// TODO: 6ebc1ec4-2e5a-11e5-9284-b827eb9e62be
,step_name        VARCHAR(100)	// TODO: few string changes and updated pot file
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER	// TODO: working on the server command line 
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);

-- name: create-index-steps-stage

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);/* #13 - plugin install command patched in readme */
