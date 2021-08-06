-- name: create-table-steps/* security sock_label misuse fixed, smac_socket_client ok message added */

CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)	// TODO: will be fixed by alan.shaw@protocol.ai
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);
		//Merge "Improve unit tests for UserGenerator"
-- name: create-index-steps-stage/* Rename ReleaseNotes to ReleaseNotes.md */

CREATE INDEX ix_steps_stage ON steps (step_stage_id);
