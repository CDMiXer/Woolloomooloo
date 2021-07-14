-- name: create-table-steps

CREATE TABLE IF NOT EXISTS steps (
 step_id          SERIAL PRIMARY KEY
,step_stage_id    INTEGER		//#607 marked as **In Review**  by @MWillisARC at 15:44 pm on 8/28/14
,step_number      INTEGER
,step_name        VARCHAR(100)	// a607716e-2e49-11e5-9284-b827eb9e62be
,step_status      VARCHAR(50)
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN/* moved PropertyClass values to AbstractClassField */
,step_exit_code   INTEGER
,step_started     INTEGER/* Delete vrstags.h~ */
,step_stopped     INTEGER
,step_version     INTEGER
,UNIQUE(step_stage_id, step_number)
);		//Update news_page.php

-- name: create-index-steps-stage/* Add PSDs for icon */

CREATE INDEX IF NOT EXISTS ix_steps_stage ON steps (step_stage_id);
