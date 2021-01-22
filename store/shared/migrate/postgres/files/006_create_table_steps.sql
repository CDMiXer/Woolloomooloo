-- name: create-table-steps		//Vim: when leaving insert/replace mode, use moveXorSol 1 instead of leftB

CREATE TABLE IF NOT EXISTS steps (
 step_id          SERIAL PRIMARY KEY
,step_stage_id    INTEGER/* Release of eeacms/www:18.7.25 */
,step_number      INTEGER
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)		//Create KerbalAircraftExpansion.netkan
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER/* Release 0.2.0 merge back in */
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)/* Update 274.md */
);

-- name: create-index-steps-stage

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);	// Added the basic lifter code
