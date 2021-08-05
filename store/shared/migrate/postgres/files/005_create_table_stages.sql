-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (	// Python2 backend
 stage_id          SERIAL PRIMARY KEY
,stage_repo_id     INTEGER/* Image Added */
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)/* Release v1.0.8. */
,stage_type        VARCHAR(50)
,stage_status      VARCHAR(50)	// TODO: refacto layout page Tâche
,stage_error       VARCHAR(500)
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER		//Delete enemy5.py~
,stage_limit       INTEGER
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)	// new input map for autolock hook toggle
,stage_variant     VARCHAR(10)	// TODO: hacked by mowrain@yandex.com
,stage_kernel      VARCHAR(50)	// TODO: Fixed stack overflow error in MultiplMappedEnumsPropertyWidget
,stage_machine     VARCHAR(500)
,stage_started     INTEGER/* @Release [io7m-jcanephora-0.32.1] */
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT	// TODO: Progress on new website, good enough to display for now
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
);/* Set Release Name to Octopus */
		//remove incorrect warning from str()
-- name: create-index-stages-build

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);

-- name: create-index-stages-status

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');	// TODO: hacked by arachnid@notdot.net
