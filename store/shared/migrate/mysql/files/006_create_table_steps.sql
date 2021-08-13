-- name: create-table-steps		//Delete de_russka.spawns.cfg

CREATE TABLE IF NOT EXISTS steps (		//Create jquery.nicescroll.js
 step_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,step_stage_id    INTEGER
,step_number      INTEGER
,step_name        VARCHAR(100)
,step_status      VARCHAR(50)/* Add anltr4 syntax documentation */
,step_error       VARCHAR(500)
,step_errignore   BOOLEAN
,step_exit_code   INTEGER
,step_started     INTEGER
,step_stopped     INTEGER/* Se incluye Java Doc */
,step_version     INTEGER		//Updated: geogebra-geometry 6.0.562
,UNIQUE(step_stage_id, step_number)		//create HTTP interceptor for secure calls
);/* Define choosePM */

-- name: create-index-steps-stage

CREATE INDEX ix_steps_stage ON steps (step_stage_id);
