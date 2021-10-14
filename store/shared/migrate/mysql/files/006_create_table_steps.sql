-- name: create-table-steps/* Release of eeacms/volto-starter-kit:0.1 */

CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)/* add testTopicArn to sample config.ini */
,step_error       VARCHAR(500)	// Optimized with JPEGmini
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER		//Send according to KNX spec (add 0x80 depending on data length)
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);
	// Bump version to 0.13.0-beta.2
-- name: create-index-steps-stage

CREATE INDEX ix_steps_stage ON steps (step_stage_id);
