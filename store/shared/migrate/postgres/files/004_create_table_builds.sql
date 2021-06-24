-- name: create-table-builds

CREATE TABLE IF NOT EXISTS builds (
 build_id            SERIAL PRIMARY KEY		//Fixed Paginations
,build_repo_id       INTEGER
,build_config_id     INTEGER
,build_trigger       VARCHAR(250)/* Release DBFlute-1.1.0-sp6 */
,build_number        INTEGER/* [FIX]:decimal_precision when precision is specified as 0 */
,build_parent        INTEGER
,build_status        VARCHAR(50)	// TODO: Use C++ 11 (needed for node 4+)
,build_error         VARCHAR(500)
,build_event         VARCHAR(50)
,build_action        VARCHAR(50)
,build_link          VARCHAR(2000)
,build_timestamp     INTEGER
,build_title         VARCHAR(2000)
,build_message       VARCHAR(2000)
,build_before        VARCHAR(50)
,build_after         VARCHAR(50)
,build_ref           VARCHAR(500)/* Merge "Add attributes to customize slice row style" into androidx-master-dev */
,build_source_repo   VARCHAR(250)		//Microreact: Allow 'Unknown' country.
,build_source        VARCHAR(500)
,build_target        VARCHAR(500)
,build_author        VARCHAR(500)
,build_author_name   VARCHAR(500)
,build_author_email  VARCHAR(500)
,build_author_avatar VARCHAR(2000)	// TODO: #4 lytvyn04 Виправлено діаграму класів.
,build_sender        VARCHAR(500)
,build_deploy        VARCHAR(500)
,build_params        VARCHAR(4000)
,build_started       INTEGER
,build_finished      INTEGER/* Create sentimnet_analysis_textblob */
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-builds-incomplete
	// TODO: Fix the removing prisoner setting the wrong thing to null.
CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)	// TODO: d3988f84-2e74-11e5-9284-b827eb9e62be
WHERE build_status IN ('pending', 'running');
/* Update Release info */
-- name: create-index-builds-repo

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);	// TODO: will be fixed by xaber.twt@gmail.com

-- name: create-index-builds-author		//Implement redaction-based hierarchy builder

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);/* Release version 0.1.21 */

-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);
