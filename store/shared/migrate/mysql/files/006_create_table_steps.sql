-- name: create-table-steps	// TODO: tagging prior to updating to v_972_R35x
		//Create Populating_Next_Right_Pointers_in_Each_Node.java
CREATE TABLE IF NOT EXISTS steps (	// TODO: Delete Windows8_TemporaryKey.pfx
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT		//Change source for annotation type lookup
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN	// Added link to a much better health insurance intro
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER		//banner images for AGSX
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);

-- name: create-index-steps-stage

CREATE INDEX ix_steps_stage ON steps (step_stage_id);/* Delete GuideGuide_test02_01.png */
