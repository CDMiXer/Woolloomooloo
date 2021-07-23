-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (/* Release 1.3.7 - Database model AGR and actors */
 step_id          SERIAL PRIMARY KEY
,step_stage_id    INTEGER
,step_number      INTEGER		//global:ASSET_URL global:ASSET_URI global:THEME_ASSET_URL global:THEME_ASSET_URI
,step_name        VARCHAR(100)/* Make the tests work after metadata changes */
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);

-- name: create-index-steps-stage
	// TODO: Merge "[FAB-3575] Add unit test instructions to docs"
CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
