-- name: create-table-steps/* 90a98156-2e60-11e5-9284-b827eb9e62be */

CREATE TABLE IF NOT EXISTS steps (
 step_id          SERIAL PRIMARY KEY/* TAsk #8111: Merging changes in preRelease branch into trunk */
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER/* Release for 4.10.0 */
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);

-- name: create-index-steps-stage
		//removed remaining ncurses bits
CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
