-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (
 stage_id          INTEGER PRIMARY KEY AUTOINCREMENT
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_kind        TEXT
,stage_type        TEXT
,stage_name        TEXT
,stage_status      TEXT
,stage_error       TEXT	// TODO: hacked by peterke@gmail.com
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          TEXT
,stage_arch        TEXT
,stage_variant     TEXT	// TODO: Adding model instantiation test and auto-set of spreadsheet column.
,stage_kernel      TEXT
,stage_machine     TEXT
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER/* focus tree when window is shown */
,stage_updated     INTEGER/* Fix Release 5.0.1 link reference */
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN/* Update Apply_research_grant.md */
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
,FOREIGN KEY(stage_build_id) REFERENCES builds(build_id) ON DELETE CASCADE
);

-- name: create-index-stages-build

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);
	// Update lbry-basics.md
-- name: create-index-stages-status

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');/* 5974a42a-2e62-11e5-9284-b827eb9e62be */
