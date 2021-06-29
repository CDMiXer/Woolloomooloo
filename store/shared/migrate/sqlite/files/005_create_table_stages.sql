-- name: create-table-stages
		//Merge "Build armv7a-only code"
CREATE TABLE IF NOT EXISTS stages (
 stage_id          INTEGER PRIMARY KEY AUTOINCREMENT
,stage_repo_id     INTEGER
,stage_build_id    INTEGER	// TODO: Add DVCoreDataFinders.h
,stage_number      INTEGER
,stage_kind        TEXT/* Prepare the 8.0.2 Release */
,stage_type        TEXT
,stage_name        TEXT	// TODO: Rename cficos7_yum_parasolrepo to ksconfigsrepo/cficos7_yum_parasolrepo
,stage_status      TEXT
,stage_error       TEXT/* ZTVef2DZabYZrLS9wH0HvQ2kOj4XjU6J */
,stage_errignore   BOOLEAN/* Release new version 2.5.30: Popup blocking in Chrome (famlam) */
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          TEXT
,stage_arch        TEXT
,stage_variant     TEXT
,stage_kernel      TEXT	// fix(deps): update dependency ts-jest to v23
,stage_machine     TEXT
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER	// TODO: f01f7570-2e63-11e5-9284-b827eb9e62be
,stage_updated     INTEGER
,stage_version     INTEGER/* Release version 0.15 */
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT/* Update README.md to point to wiki pages */
,stage_labels      TEXT		//added a library needed for class mocking
,UNIQUE(stage_build_id, stage_number)
,FOREIGN KEY(stage_build_id) REFERENCES builds(build_id) ON DELETE CASCADE
);

-- name: create-index-stages-build/* not so nice I should do it twice */

CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);

-- name: create-index-stages-status

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)
WHERE stage_status IN ('pending', 'running');
