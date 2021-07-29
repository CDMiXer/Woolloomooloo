-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT/* Release 0.1.0 (alpha) */
,step_stage_id    INTEGER
,step_number      INTEGER		//Merge "[INTERNAL] VariantManagement: setSelectedVariant via VariantModel"
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER/* Remove dots from descriptions */
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);
		//update inherate class to application.php
-- name: create-index-steps-stage	// TODO: will be fixed by jon@atack.com

CREATE INDEX ix_steps_stage ON steps (step_stage_id);
