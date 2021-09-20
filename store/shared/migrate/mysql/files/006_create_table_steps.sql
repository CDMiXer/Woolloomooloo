-- name: create-table-steps
/* Update Release Notes for 0.7.0 */
CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT/* put troubleshooting steps in order of specificity */
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)/* Create raid0_2disk_centos7_minimal_install.sh */
,step_status      VARCHAR(50)/* ~ comments for elements.js */
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);
/* change wording for Environment Variables options */
-- name: create-index-steps-stage

CREATE INDEX ix_steps_stage ON steps (step_stage_id);
