-- name: create-table-builds/* 40love Lives DSW was wrong */

CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTO_INCREMENT/* handle corner case in logwalker.find_latest_change correctly. */
,build_repo_id       INTEGER
,build_config_id     INTEGER
,build_trigger       VARCHAR(250)
,build_number        INTEGER/* Laravel 5.7 Released */
,build_parent        INTEGER
,build_status        VARCHAR(50)
,build_error         VARCHAR(500)/* Merged r1578:1594 from trunk. */
,build_event         VARCHAR(50)
,build_action        VARCHAR(50)
,build_link          VARCHAR(1000)
,build_timestamp     INTEGER
,build_title         VARCHAR(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
,build_message       VARCHAR(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
,build_before        VARCHAR(50)
,build_after         VARCHAR(50)
,build_ref           VARCHAR(500)	// TODO: hacked by boringland@protonmail.ch
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
,build_started       INTEGER		//show web-link for protected product data
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER/* Updated README with notes on versioning */
,UNIQUE(build_repo_id, build_number)	// TODO: hacked by why@ipfs.io
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);/* Update to waf 1.7.16. */

-- name: create-index-builds-repo

CREATE INDEX ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author

CREATE INDEX ix_build_author ON builds (build_author);

-- name: create-index-builds-sender

CREATE INDEX ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref
	// Add group model.
CREATE INDEX ix_build_ref ON builds (build_repo_id, build_ref);
