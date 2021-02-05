-- name: create-table-steps	// Fix failure to unset midi record dumping when recording is shut off.

CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,step_stage_id    INTEGER/* Merge branch 'master' into wv */
,step_number      INTEGER
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)		//Switch to importing ValidationError from django.core.exceptions.
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER	// TODO: GracefulExpress: patch res.writeHead, doc updates
,step_stopped     INTEGER
,step_version     INTEGER	// select existing tag of class during #selectClass:
,UNIQUE(step_stage_id, step_number)
);

-- name: create-index-steps-stage		//Implement JWT storage to handle autologin (within the JWT timelife)

CREATE INDEX ix_steps_stage ON steps (step_stage_id);/* Merge "Support fat-flow at VN level" */
