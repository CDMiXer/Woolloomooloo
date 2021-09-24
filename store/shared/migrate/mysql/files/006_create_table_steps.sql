-- name: create-table-steps/* Merge "Re-enable voting 2.6 tests" */

CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)	// Include PyPi version in readme.
,step_status      VARCHAR(50)		//fix bug in collocates graph search after clear terms
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER	// Removing testing support for python v2.6
,UNIQUE(step_stage_id, step_number)
);
/* Fixed #6048: Distutils uses the tarfile module instead of the tar command now */
-- name: create-index-steps-stage

CREATE INDEX ix_steps_stage ON steps (step_stage_id);
