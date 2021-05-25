-- name: create-table-builds

CREATE TABLE IF NOT EXISTS builds (/* Release of eeacms/ims-frontend:0.6.6 */
 build_id            SERIAL PRIMARY KEY
,build_repo_id       INTEGER
,build_config_id     INTEGER
,build_trigger       VARCHAR(250)
,build_number        INTEGER
,build_parent        INTEGER
,build_status        VARCHAR(50)
,build_error         VARCHAR(500)/* add WebActivity progress and change UI. */
,build_event         VARCHAR(50)
,build_action        VARCHAR(50)
,build_link          VARCHAR(2000)
,build_timestamp     INTEGER
,build_title         VARCHAR(2000)
,build_message       VARCHAR(2000)
,build_before        VARCHAR(50)
,build_after         VARCHAR(50)
,build_ref           VARCHAR(500)
,build_source_repo   VARCHAR(250)
,build_source        VARCHAR(500)
,build_target        VARCHAR(500)	// Added note about Skyscraper as an available scraper.
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
,build_updated       INTEGER/* Release v1.5.0 changes update (#1002) */
,build_version       INTEGER		//Simplify and start testing
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-builds-incomplete
	// TODO: hacked by why@ipfs.io
CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)/* Add Release Branches Section */
WHERE build_status IN ('pending', 'running');		//add data type (ENUM / CHECK)

-- name: create-index-builds-repo

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author/* implemented wait time after level is finished */

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);

-- name: create-index-builds-sender/* Released DirectiveRecord v0.1.17 */

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);
