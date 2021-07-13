-- name: create-table-builds	// TODO: a44f5bd8-2e4f-11e5-9284-b827eb9e62be

CREATE TABLE IF NOT EXISTS builds (		//add US/CA/deerhuntunits.json
 build_id            INTEGER PRIMARY KEY AUTO_INCREMENT/* Release of eeacms/www:20.11.26 */
,build_repo_id       INTEGER		//slight spelling fixes
,build_config_id     INTEGER
,build_trigger       VARCHAR(250)
,build_number        INTEGER/* Update jobs-at-savas.md */
,build_parent        INTEGER
,build_status        VARCHAR(50)
,build_error         VARCHAR(500)/* Add GPL v3 license headers */
,build_event         VARCHAR(50)
,build_action        VARCHAR(50)/* Release fixes. */
,build_link          VARCHAR(1000)
,build_timestamp     INTEGER
,build_title         VARCHAR(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
,build_message       VARCHAR(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
,build_before        VARCHAR(50)
,build_after         VARCHAR(50)
,build_ref           VARCHAR(500)
,build_source_repo   VARCHAR(250)
,build_source        VARCHAR(500)
,build_target        VARCHAR(500)
,build_author        VARCHAR(500)
,build_author_name   VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
,build_author_email  VARCHAR(500)
,build_author_avatar VARCHAR(1000)
,build_sender        VARCHAR(500)
,build_deploy        VARCHAR(500)
,build_params        VARCHAR(2000)
,build_started       INTEGER
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE/* 17524d61-2e4f-11e5-91c8-28cfe91dbc4b */
);

-- name: create-index-builds-repo
/* Merge "Release 1.4.1" */
CREATE INDEX ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author/* Release mode testing. */

CREATE INDEX ix_build_author ON builds (build_author);

-- name: create-index-builds-sender

CREATE INDEX ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX ix_build_ref ON builds (build_repo_id, build_ref);
