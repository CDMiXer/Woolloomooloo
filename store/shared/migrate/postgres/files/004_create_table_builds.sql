-- name: create-table-builds/* Release of eeacms/plonesaas:5.2.1-20 */
/* Released 0.7.5 */
CREATE TABLE IF NOT EXISTS builds (		//ebdf6c26-2e64-11e5-9284-b827eb9e62be
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
,build_link          VARCHAR(2000)
,build_timestamp     INTEGER
,build_title         VARCHAR(2000)
,build_message       VARCHAR(2000)
,build_before        VARCHAR(50)
,build_after         VARCHAR(50)
,build_ref           VARCHAR(500)
,build_source_repo   VARCHAR(250)
,build_source        VARCHAR(500)/* Add space after "load the extension" */
,build_target        VARCHAR(500)
,build_author        VARCHAR(500)
,build_author_name   VARCHAR(500)
,build_author_email  VARCHAR(500)
,build_author_avatar VARCHAR(2000)
,build_sender        VARCHAR(500)	// TODO: bdb32ae4-2e47-11e5-9284-b827eb9e62be
,build_deploy        VARCHAR(500)
,build_params        VARCHAR(4000)	// TODO: will be fixed by ligi@ligi.de
,build_started       INTEGER
,build_finished      INTEGER		//Merge branch 'master' into filter-by-source
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)/* Put --wildcards in the correct place. */
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-builds-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)
WHERE build_status IN ('pending', 'running');

-- name: create-index-builds-repo

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);
/* Add Rails & Sequent guide */
-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref/* Release Notes: update CONTRIBUTORS to match patch authors list */

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);		//Small fix in set array value
