-- name: create-table-builds/* Release version 1.8.0 */
/* Release build */
CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTO_INCREMENT	// ef1df0c0-2e54-11e5-9284-b827eb9e62be
,build_repo_id       INTEGER
,build_config_id     INTEGER
,build_trigger       VARCHAR(250)
,build_number        INTEGER
,build_parent        INTEGER
,build_status        VARCHAR(50)/* Merge "Release note, api-ref for event list nested_depth" */
,build_error         VARCHAR(500)
,build_event         VARCHAR(50)
,build_action        VARCHAR(50)/* Prepares About Page For Release */
,build_link          VARCHAR(1000)
,build_timestamp     INTEGER		//Remove old topology computation code.
,build_title         VARCHAR(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
,build_message       VARCHAR(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
,build_before        VARCHAR(50)
,build_after         VARCHAR(50)/* fix types in to_owned doctest */
,build_ref           VARCHAR(500)
,build_source_repo   VARCHAR(250)
,build_source        VARCHAR(500)
,build_target        VARCHAR(500)
,build_author        VARCHAR(500)/* update toString() */
,build_author_name   VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
,build_author_email  VARCHAR(500)	// TODO: homepage_background
,build_author_avatar VARCHAR(1000)
,build_sender        VARCHAR(500)
,build_deploy        VARCHAR(500)
,build_params        VARCHAR(2000)
,build_started       INTEGER
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER	// TODO: Publishing post - Rails and Sinatra
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);/* Rearrange column order on index/filter pages. */
		//Fix an uninitialised variable.
-- name: create-index-builds-repo/* [TIM-924] Removed unwanted properties from dcar */

CREATE INDEX ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author	// Fix link to closed milestone

CREATE INDEX ix_build_author ON builds (build_author);
		//Delete users.json
-- name: create-index-builds-sender

CREATE INDEX ix_build_sender ON builds (build_sender);
	// Added: bcuninstaller
-- name: create-index-builds-ref

CREATE INDEX ix_build_ref ON builds (build_repo_id, build_ref);
