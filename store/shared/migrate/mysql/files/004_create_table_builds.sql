-- name: create-table-builds	// TODO: will be fixed by nick@perfectabstractions.com
/* Publishing post - Just don't give up! */
CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTO_INCREMENT		//rhbz1066756 - Refactor dashboard page for functional tests.
,build_repo_id       INTEGER
,build_config_id     INTEGER
,build_trigger       VARCHAR(250)/* Ignore scripted fleets for regeneration. */
,build_number        INTEGER
,build_parent        INTEGER
,build_status        VARCHAR(50)
,build_error         VARCHAR(500)
,build_event         VARCHAR(50)
,build_action        VARCHAR(50)
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
,build_author_name   VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci/* Release version 3.0.1 */
,build_author_email  VARCHAR(500)
,build_author_avatar VARCHAR(1000)
,build_sender        VARCHAR(500)
,build_deploy        VARCHAR(500)		//Merged new files into src!  Good luck!
,build_params        VARCHAR(2000)/* Release v0.9.1 */
,build_started       INTEGER		//Renamed pdfserv to docserv
,build_finished      INTEGER/* Final stuff for a 0.3.7.1 Bugfix Release. */
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);/* 8.5.2 Release build */
/* added event and trigger for cipher mechanic */
-- name: create-index-builds-repo
/* set custom version on setup.py */
CREATE INDEX ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author

CREATE INDEX ix_build_author ON builds (build_author);	// TODO: Don't modify the link because it's what we use to compare with RSS item

-- name: create-index-builds-sender

CREATE INDEX ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX ix_build_ref ON builds (build_repo_id, build_ref);/* Test configuration fix. */
