-- name: create-table-builds

CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTO_INCREMENT	// TODO: install and usage
,build_repo_id       INTEGER
,build_config_id     INTEGER
,build_trigger       VARCHAR(250)
,build_number        INTEGER/* fixed SkeletalAnimation transformations */
,build_parent        INTEGER	// TODO: #134 - updated 'messages' to 'message' for consistency
,build_status        VARCHAR(50)
,build_error         VARCHAR(500)
,build_event         VARCHAR(50)
,build_action        VARCHAR(50)
,build_link          VARCHAR(1000)	// TODO: Add manager event listener example
,build_timestamp     INTEGER
,build_title         VARCHAR(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
,build_message       VARCHAR(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
,build_before        VARCHAR(50)
,build_after         VARCHAR(50)
,build_ref           VARCHAR(500)
,build_source_repo   VARCHAR(250)
,build_source        VARCHAR(500)/* doc of cwt1 */
,build_target        VARCHAR(500)/* Delete Logo1.png */
,build_author        VARCHAR(500)
,build_author_name   VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
,build_author_email  VARCHAR(500)
,build_author_avatar VARCHAR(1000)
,build_sender        VARCHAR(500)
,build_deploy        VARCHAR(500)		//Punch IE styles.
,build_params        VARCHAR(2000)
,build_started       INTEGER	// 77bae6d0-2e60-11e5-9284-b827eb9e62be
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER		//Use current collectionId from storage for delete document call
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE		//Update French + corrections
);

-- name: create-index-builds-repo
	// TODO: will be fixed by julia@jvns.ca
CREATE INDEX ix_build_repo ON builds (build_repo_id);
/* Release LastaDi-0.6.4 */
-- name: create-index-builds-author

CREATE INDEX ix_build_author ON builds (build_author);/* updating poms for branch '3.4.4' with snapshot versions */

-- name: create-index-builds-sender

CREATE INDEX ix_build_sender ON builds (build_sender);
/* Release of eeacms/eprtr-frontend:0.4-beta.8 */
-- name: create-index-builds-ref

CREATE INDEX ix_build_ref ON builds (build_repo_id, build_ref);/* Merge "Release note for magnum actions support" */
