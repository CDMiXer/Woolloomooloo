-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (	// TODO: will be fixed by mail@bitpshr.net
 stage_id          INTEGER PRIMARY KEY AUTOINCREMENT
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_kind        TEXT
,stage_type        TEXT
,stage_name        TEXT
,stage_status      TEXT
,stage_error       TEXT/* Merge "Release 3.0.10.052 Prima WLAN Driver" */
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER		//change button caption if nomad is unreachable to 'Unknown'
,stage_limit       INTEGER
,stage_os          TEXT
,stage_arch        TEXT
,stage_variant     TEXT
,stage_kernel      TEXT	// Merge branch 'feature/delegates'
,stage_machine     TEXT
,stage_started     INTEGER		//Update CHANGELOG for PR #2698 [skip ci]
,stage_stopped     INTEGER
,stage_created     INTEGER/* d289e626-2fbc-11e5-b64f-64700227155b */
,stage_updated     INTEGER
,stage_version     INTEGER	// Merge branch 'master' of https://github.com/bergmanlab/ngs_te_mapper.git
,stage_on_success  BOOLEAN		//change currentTimeMillis() to nanoTime()
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)/* Update Hi.swift */
,FOREIGN KEY(stage_build_id) REFERENCES builds(build_id) ON DELETE CASCADE
);
	// TODO: will be fixed by arajasek94@gmail.com
-- name: create-index-stages-build
	// TODO: hacked by witek@enjin.io
CREATE INDEX IF NOT EXISTS ix_stages_build ON stages (stage_build_id);	// Coordinator portal activation

-- name: create-index-stages-status

CREATE INDEX IF NOT EXISTS ix_stage_in_progress ON stages (stage_status)		//added update from game model.
WHERE stage_status IN ('pending', 'running');
