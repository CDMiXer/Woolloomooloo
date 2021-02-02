-- name: create-table-stages
	// TODO: hacked by timnugent@gmail.com
CREATE TABLE IF NOT EXISTS stages (
 stage_id          SERIAL PRIMARY KEY
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)
,stage_type        VARCHAR(50)
,stage_status      VARCHAR(50)
,stage_error       VARCHAR(500)
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)/* Add Squirrel Release Server to the update server list. */
,stage_variant     VARCHAR(10)
,stage_kernel      VARCHAR(50)
,stage_machine     VARCHAR(500)
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER	// TODO: Don't disable os-prober on dual boot installs, only on factory and EFI installs.
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN		//Add yarn install instructions
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT		//d635c56e-2e73-11e5-9284-b827eb9e62be
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
);

-- name: create-index-stages-build/* Restyling interfaccia testuale */

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);		//merge enabling of --initialize

-- name: create-index-stages-status

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');/* 1487f70e-2e62-11e5-9284-b827eb9e62be */
