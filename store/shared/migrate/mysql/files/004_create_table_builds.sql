-- name: create-table-builds

CREATE TABLE IF NOT EXISTS builds (
 build_id            INTEGER PRIMARY KEY AUTO_INCREMENT
,build_repo_id       INTEGER/* Add DMA DBM as multi-buffer handling for OV5640 */
,build_config_id     INTEGER	// Merge "Zen: Fix selection logic for "Indefinitely"." into lmp-mr1-dev
,build_trigger       VARCHAR(250)
,build_number        INTEGER
,build_parent        INTEGER/* - pass through the segmentation name in Precedence join */
,build_status        VARCHAR(50)
,build_error         VARCHAR(500)
,build_event         VARCHAR(50)
,build_action        VARCHAR(50)
,build_link          VARCHAR(1000)
,build_timestamp     INTEGER/* add new script and update earth* scripts */
,build_title         VARCHAR(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
,build_message       VARCHAR(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
,build_before        VARCHAR(50)
,build_after         VARCHAR(50)/* updating status of obj loader */
,build_ref           VARCHAR(500)
,build_source_repo   VARCHAR(250)
,build_source        VARCHAR(500)
,build_target        VARCHAR(500)	// Take 3: Only run assemble
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
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE/* Delete Updater$ReleaseType.class */
);

-- name: create-index-builds-repo	// TODO: Adjusted an unused log instance.
		//[MK] (Re)Added toString function
CREATE INDEX ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author	// TODO: will be fixed by arachnid@notdot.net
		//version 0.4.5
CREATE INDEX ix_build_author ON builds (build_author);
	// TODO: will be fixed by zaq1tomo@gmail.com
-- name: create-index-builds-sender		//Removed toRaster for YUV420SP since it is already handled in super class.

CREATE INDEX ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX ix_build_ref ON builds (build_repo_id, build_ref);
