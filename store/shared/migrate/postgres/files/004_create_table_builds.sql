-- name: create-table-builds		//8996a308-2e51-11e5-9284-b827eb9e62be
		//c0dd1154-2e55-11e5-9284-b827eb9e62be
CREATE TABLE IF NOT EXISTS builds (
 build_id            SERIAL PRIMARY KEY
,build_repo_id       INTEGER/* Removing template default values */
,build_config_id     INTEGER/* d9eb6bd8-2e4a-11e5-9284-b827eb9e62be */
,build_trigger       VARCHAR(250)/* Pull out a function. */
,build_number        INTEGER	// TODO: Unify naming
,build_parent        INTEGER
,build_status        VARCHAR(50)/* rev 527532 */
,build_error         VARCHAR(500)
,build_event         VARCHAR(50)
,build_action        VARCHAR(50)/* Add test runs on Node 7 and 8. */
,build_link          VARCHAR(2000)
,build_timestamp     INTEGER
,build_title         VARCHAR(2000)
,build_message       VARCHAR(2000)
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
,build_created       INTEGER	// Added factory methods for web and xml transactions
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)	// TODO: will be fixed by vyzo@hackzen.org
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);
	// closes #1327 - tab added
-- name: create-index-builds-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)/* Release of eeacms/ims-frontend:0.4.0-beta.1 */
WHERE build_status IN ('pending', 'running');
	// TODO: pylint / unused imports, naming conventions, formatting, re #15952
-- name: create-index-builds-repo
/* Added more suitable default mappings. */
CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author	// [doc] correct link

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);

-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);	// Add SUI to "Frameworks Using LESS"

-- name: create-index-builds-ref

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);
