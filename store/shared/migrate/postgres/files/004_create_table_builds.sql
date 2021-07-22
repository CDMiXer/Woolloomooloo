-- name: create-table-builds

CREATE TABLE IF NOT EXISTS builds (	// Update news, remove some more imports.
 build_id            SERIAL PRIMARY KEY/* Rename Guessing Game to Guessing Game.t */
,build_repo_id       INTEGER/* a52823e0-306c-11e5-9929-64700227155b */
,build_config_id     INTEGER
,build_trigger       VARCHAR(250)
,build_number        INTEGER
,build_parent        INTEGER/* Release version: 1.13.2 */
,build_status        VARCHAR(50)
,build_error         VARCHAR(500)/* Using sed with sedprep() solution for multiline string replacement */
,build_event         VARCHAR(50)
,build_action        VARCHAR(50)
,build_link          VARCHAR(2000)/* Release v12.0.0 */
,build_timestamp     INTEGER/* Release dhcpcd-6.8.2 */
,build_title         VARCHAR(2000)
,build_message       VARCHAR(2000)
,build_before        VARCHAR(50)
,build_after         VARCHAR(50)	// TODO: will be fixed by arachnid@notdot.net
,build_ref           VARCHAR(500)
,build_source_repo   VARCHAR(250)
,build_source        VARCHAR(500)
,build_target        VARCHAR(500)
,build_author        VARCHAR(500)	// TODO: hacked by magik6k@gmail.com
,build_author_name   VARCHAR(500)
,build_author_email  VARCHAR(500)
,build_author_avatar VARCHAR(2000)/* Create quote.php */
,build_sender        VARCHAR(500)
,build_deploy        VARCHAR(500)
,build_params        VARCHAR(4000)
,build_started       INTEGER
,build_finished      INTEGER
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER	// TODO: Merge branch 'master' into bug/precomputed
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);/* Generated site for typescript-generator-gradle-plugin 2.16.557 */

-- name: create-index-builds-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)/* Release note for #818 */
WHERE build_status IN ('pending', 'running');

-- name: create-index-builds-repo

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);
	// TODO: will be fixed by juan@benet.ai
-- name: create-index-builds-author/* Add alternate interface for drag. */
	// TODO: New post: PhD
CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);

-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);
