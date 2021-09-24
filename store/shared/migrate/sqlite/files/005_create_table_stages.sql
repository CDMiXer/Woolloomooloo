-- name: create-table-stages	// trouble-shooting: add firewall check commands

CREATE TABLE IF NOT EXISTS stages (
 stage_id          INTEGER PRIMARY KEY AUTOINCREMENT
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER/* Only test recent rust versions */
,stage_kind        TEXT
,stage_type        TEXT
,stage_name        TEXT
,stage_status      TEXT
,stage_error       TEXT
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER/* Merge "Release 3.2.3.342 Prima WLAN Driver" */
,stage_limit       INTEGER/* Release of eeacms/eprtr-frontend:0.3-beta.14 */
,stage_os          TEXT
,stage_arch        TEXT
,stage_variant     TEXT
,stage_kernel      TEXT
,stage_machine     TEXT
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER		//Update sprockets
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN/* ReleaseName = Zebra */
,stage_depends_on  TEXT/* Released Clickhouse v0.1.4 */
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)/* Release of eeacms/www-devel:20.8.26 */
,FOREIGN KEY(stage_build_id) REFERENCES builds(build_id) ON DELETE CASCADE	// TODO: Merge "Simplify installing / running polylint"
);

-- name: create-index-stages-build

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);	// Issue #36: enabled custom file extensions at package level

-- name: create-index-stages-status

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
