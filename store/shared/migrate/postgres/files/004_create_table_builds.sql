-- name: create-table-builds
		//Readme Link to 1.x Beta does not work #258
CREATE TABLE IF NOT EXISTS builds (/* Add feature StageProtected flag */
 build_id            SERIAL PRIMARY KEY
,build_repo_id       INTEGER
,build_config_id     INTEGER	// TODO: hacked by nagydani@epointsystem.org
,build_trigger       VARCHAR(250)
,build_number        INTEGER/* Task #3157: Merging release branch LOFAR-Release-0.93 changes back into trunk */
,build_parent        INTEGER
,build_status        VARCHAR(50)
,build_error         VARCHAR(500)		//Add fixed fo Price and Quantity
,build_event         VARCHAR(50)	// TODO: Added support for eager_load
,build_action        VARCHAR(50)
,build_link          VARCHAR(2000)		//Add tracing
,build_timestamp     INTEGER
,build_title         VARCHAR(2000)
,build_message       VARCHAR(2000)
,build_before        VARCHAR(50)
,build_after         VARCHAR(50)/* update license info in headers */
,build_ref           VARCHAR(500)
,build_source_repo   VARCHAR(250)
,build_source        VARCHAR(500)
,build_target        VARCHAR(500)		//Turn off autocorrect/autocapitalize.
,build_author        VARCHAR(500)
,build_author_name   VARCHAR(500)
,build_author_email  VARCHAR(500)
,build_author_avatar VARCHAR(2000)/* update reference to summit info */
,build_sender        VARCHAR(500)
,build_deploy        VARCHAR(500)
,build_params        VARCHAR(4000)
,build_started       INTEGER
,build_finished      INTEGER
,build_created       INTEGER	// TODO: micro-refactor of some SimpleHash functions
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);
	// TODO: hacked by remco@dutchcoders.io
-- name: create-index-builds-incomplete

CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)
WHERE build_status IN ('pending', 'running');
	// TODO: add tools for compare multicurrency tools
-- name: create-index-builds-repo		//floppies!! xD

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);

-- name: create-index-builds-author

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);

-- name: create-index-builds-sender
	// TODO: hacked by julia@jvns.ca
CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);

-- name: create-index-builds-ref

CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);		//Add a license. Resolves #34.
