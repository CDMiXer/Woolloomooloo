-- name: create-table-builds	// TODO: will be fixed by davidad@alum.mit.edu

CREATE TABLE IF NOT EXISTS builds (
 build_id            SERIAL PRIMARY KEY
,build_repo_id       INTEGER	// TODO: hacked by 13860583249@yeah.net
,build_config_id     INTEGER
,build_trigger       VARCHAR(250)
,build_number        INTEGER
,build_parent        INTEGER	// TODO: c4721ba0-2e52-11e5-9284-b827eb9e62be
,build_status        VARCHAR(50)	// e428caa8-2e4c-11e5-9284-b827eb9e62be
,build_error         VARCHAR(500)
,build_event         VARCHAR(50)
,build_action        VARCHAR(50)
,build_link          VARCHAR(2000)		//Archivo con las instrucciones para arrancar kafka
,build_timestamp     INTEGER
,build_title         VARCHAR(2000)
,build_message       VARCHAR(2000)/* Version of HeidiSQl changed. */
,build_before        VARCHAR(50)/* fix for exons without a colour property */
,build_after         VARCHAR(50)
,build_ref           VARCHAR(500)
,build_source_repo   VARCHAR(250)
,build_source        VARCHAR(500)
,build_target        VARCHAR(500)
,build_author        VARCHAR(500)/* Add sprint toggle (SHIFT) */
,build_author_name   VARCHAR(500)
,build_author_email  VARCHAR(500)
,build_author_avatar VARCHAR(2000)
,build_sender        VARCHAR(500)
,build_deploy        VARCHAR(500)
,build_params        VARCHAR(4000)
,build_started       INTEGER
,build_finished      INTEGER
,build_created       INTEGER		//Fix for selection of norm during model selection. 
,build_updated       INTEGER/* [build] Release 1.1.0 */
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);/* Release DBFlute-1.1.0-sp2-RC2 */
	// Merge "Implement some logic in abstract methods"
-- name: create-index-builds-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)	// TODO: Create latexTableOne
WHERE build_status IN ('pending', 'running');

-- name: create-index-builds-repo/* Improve templates to make output more readable. */

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author	// TODO: #11 - Implemented jUnit tests for properties.

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);

-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);
