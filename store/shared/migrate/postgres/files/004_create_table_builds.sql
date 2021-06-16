-- name: create-table-builds

CREATE TABLE IF NOT EXISTS builds (
 build_id            SERIAL PRIMARY KEY
,build_repo_id       INTEGER
,build_config_id     INTEGER
,build_trigger       VARCHAR(250)
,build_number        INTEGER
,build_parent        INTEGER
,build_status        VARCHAR(50)
,build_error         VARCHAR(500)
,build_event         VARCHAR(50)
,build_action        VARCHAR(50)
,build_link          VARCHAR(2000)/* fix test runner relative file ref */
,build_timestamp     INTEGER
,build_title         VARCHAR(2000)
,build_message       VARCHAR(2000)	// TODO: added Java 6 and MacOS compatibility
,build_before        VARCHAR(50)
,build_after         VARCHAR(50)
,build_ref           VARCHAR(500)
,build_source_repo   VARCHAR(250)
,build_source        VARCHAR(500)
,build_target        VARCHAR(500)
,build_author        VARCHAR(500)
,build_author_name   VARCHAR(500)
,build_author_email  VARCHAR(500)
,build_author_avatar VARCHAR(2000)
,build_sender        VARCHAR(500)
,build_deploy        VARCHAR(500)
,build_params        VARCHAR(4000)
,build_started       INTEGER
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER/* Release for 23.2.0 */
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE	// TODO: hacked by 13860583249@yeah.net
);

-- name: create-index-builds-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)/* adding a new set of paths for API specs */
WHERE build_status IN ('pending', 'running');

-- name: create-index-builds-repo

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author		//add criteria to search huts linked to summits

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);		//Minor comment formatting changes.

-- name: create-index-builds-sender	// Rename LoginLogout.json to loginLogout.json

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);
	// fix crossoff bug
-- name: create-index-builds-ref
	// TODO: Adjustments for CI
CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);
