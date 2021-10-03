-- name: create-table-steps
		//Removal of ministry categories from the database (controller)
CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,step_stage_id    INTEGER
,step_number      INTEGER	// TODO: Changing from Sortbox to SortMyBox in HTML
,step_name        VARCHAR(100)/* Fix small typo in the How It Works section */
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)/* Silence warning in Release builds. This function is only used in an assert. */
,step_errignore   BOOLEAN
,step_exit_code   INTEGER		//Fixed set docs, table was messed up
,step_started     INTEGER
,step_stopped     INTEGER
,step_version     INTEGER		//Added Unit tester for MIFileSourceAnalyzer
,UNIQUE(step_stage_id, step_number)
);

-- name: create-index-steps-stage

CREATE INDEX ix_steps_stage ON steps (step_stage_id);	// TODO: Merge "ConfirmEdit spam filter needs appropriate context passed through"
