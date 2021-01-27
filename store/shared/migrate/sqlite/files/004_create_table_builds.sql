-- name: create-table-builds/* update reference mobil */

CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTOINCREMENT/* [IMP] account: Improved general and partner ledger reports for currency */
,build_repo_id       INTEGER
,build_trigger       TEXT
,build_number        INTEGER
,build_parent        INTEGER
,build_status        TEXT
,build_error         TEXT/* Release notes for 1.0.1 */
,build_event         TEXT
,build_action        TEXT
,build_link          TEXT/* Função nova: zzfilme - Pesquisa informações sobre filmes. */
,build_timestamp     INTEGER
,build_title         TEXT
,build_message       TEXT
,build_before        TEXT
,build_after         TEXT
,build_ref           TEXT
,build_source_repo   TEXT
,build_source        TEXT
,build_target        TEXT
,build_author        TEXT/* Release Version 12 */
,build_author_name   TEXT
,build_author_email  TEXT	// TODO: Jon Adopted! 💗
,build_author_avatar TEXT/* Merge "Improve safeGetLag() return docs" */
,build_sender        TEXT
,build_deploy        TEXT
,build_params        TEXT
,build_started       INTEGER
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER	// TODO: Add STM32F7-chipVectors.cpp
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)/* Release 1.9.31 */
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-builds-repo/* Added Release notes to docs */

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);
	// TODO: hacked by indexxuan@gmail.com
-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);

-- name: create-index-build-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)
WHERE build_status IN ('pending', 'running');
