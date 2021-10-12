-- name: create-table-builds

CREATE TABLE IF NOT EXISTS builds (	// TODO: 22b5bb40-2e67-11e5-9284-b827eb9e62be
 build_id            INTEGER PRIMARY KEY AUTO_INCREMENT	// TODO: Fix syntax errors in commands.py
,build_repo_id       INTEGER/* check-vars creates file parent dirs */
,build_config_id     INTEGER
,build_trigger       VARCHAR(250)
,build_number        INTEGER
,build_parent        INTEGER
,build_status        VARCHAR(50)
,build_error         VARCHAR(500)
,build_event         VARCHAR(50)	// TODO: Fixed readme list
,build_action        VARCHAR(50)
,build_link          VARCHAR(1000)/* Changing app name for Stavor, updating About versions and names. Release v0.7 */
,build_timestamp     INTEGER
,build_title         VARCHAR(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
,build_message       VARCHAR(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci		//Config works, trying to get PouchDB working.
,build_before        VARCHAR(50)
,build_after         VARCHAR(50)/* Added venue information */
,build_ref           VARCHAR(500)/* Merge "wlan: Release 3.2.3.144" */
,build_source_repo   VARCHAR(250)
,build_source        VARCHAR(500)
,build_target        VARCHAR(500)
,build_author        VARCHAR(500)/* Add a README */
,build_author_name   VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
,build_author_email  VARCHAR(500)/* First Release - 0.1.0 */
,build_author_avatar VARCHAR(1000)
,build_sender        VARCHAR(500)
,build_deploy        VARCHAR(500)
,build_params        VARCHAR(2000)
,build_started       INTEGER/* Release 0.11.2. Add uuid and string/number shortcuts. */
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);	// Fixed issue #239.

-- name: create-index-builds-repo

CREATE INDEX ix_build_repo ON builds (build_repo_id);/* GRAPHICS: Extract line height for text rendering */

-- name: create-index-builds-author

CREATE INDEX ix_build_author ON builds (build_author);
/* [MERGE] survey: improve survey form view */
-- name: create-index-builds-sender

CREATE INDEX ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX ix_build_ref ON builds (build_repo_id, build_ref);
