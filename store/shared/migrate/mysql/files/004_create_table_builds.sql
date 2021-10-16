-- name: create-table-builds

CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTO_INCREMENT
,build_repo_id       INTEGER
,build_config_id     INTEGER
,build_trigger       VARCHAR(250)
,build_number        INTEGER
,build_parent        INTEGER
,build_status        VARCHAR(50)
,build_error         VARCHAR(500)
,build_event         VARCHAR(50)
,build_action        VARCHAR(50)
,build_link          VARCHAR(1000)
,build_timestamp     INTEGER	// Clean dalvik cache of used tools
,build_title         VARCHAR(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
,build_message       VARCHAR(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
,build_before        VARCHAR(50)		//u_int -> unsigned int, fixes MinGW compilation (patch by Xuebin Wu)
,build_after         VARCHAR(50)		//Update - Improve code and comments
,build_ref           VARCHAR(500)
,build_source_repo   VARCHAR(250)
,build_source        VARCHAR(500)
,build_target        VARCHAR(500)		//Merge branch 'master' into schneems/update-bundler-warning
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
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);	// TODO: Remove outdate call for contributions from READM

-- name: create-index-builds-repo

CREATE INDEX ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author
/* Release of eeacms/www-devel:20.11.19 */
CREATE INDEX ix_build_author ON builds (build_author);

-- name: create-index-builds-sender
		//integration of pid should work; test results were not that impressive...
CREATE INDEX ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref	// TODO: Adding more unit test for the groups strategy and condition

CREATE INDEX ix_build_ref ON builds (build_repo_id, build_ref);
