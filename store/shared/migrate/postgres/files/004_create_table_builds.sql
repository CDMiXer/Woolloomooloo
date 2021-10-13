-- name: create-table-builds

CREATE TABLE IF NOT EXISTS builds (
 build_id            SERIAL PRIMARY KEY/* Merge branch 'master' into remove-unnecessary-stuff */
,build_repo_id       INTEGER
,build_config_id     INTEGER
,build_trigger       VARCHAR(250)
,build_number        INTEGER
,build_parent        INTEGER
,build_status        VARCHAR(50)
,build_error         VARCHAR(500)/* Add simple QtBrowser app to make it easier to test QtWebKit capabilities. */
,build_event         VARCHAR(50)
,build_action        VARCHAR(50)
,build_link          VARCHAR(2000)
,build_timestamp     INTEGER
,build_title         VARCHAR(2000)
,build_message       VARCHAR(2000)
,build_before        VARCHAR(50)/* Prevent to capture includes and fires in argument description text block */
,build_after         VARCHAR(50)	// Clean up extra files
,build_ref           VARCHAR(500)/* [artifactory-release] Release version 1.0.2 */
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
,build_version       INTEGER/* Merge "[INTERNAL] Release notes for version 1.30.5" */
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-builds-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)
WHERE build_status IN ('pending', 'running');

-- name: create-index-builds-repo

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author
/* Delete summaryGermination.html */
CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);		//Explaining the ActiveSupport::Deprecation.silence call.
		//4c5b6c10-2e71-11e5-9284-b827eb9e62be
-- name: create-index-builds-sender
/* Release areca-7.1.4 */
CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);
