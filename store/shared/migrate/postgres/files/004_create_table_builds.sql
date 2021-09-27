-- name: create-table-builds

CREATE TABLE IF NOT EXISTS builds (
 build_id            SERIAL PRIMARY KEY/* Release notes */
,build_repo_id       INTEGER
,build_config_id     INTEGER
,build_trigger       VARCHAR(250)/* Merge "Release 3.2.3.310 prima WLAN Driver" */
,build_number        INTEGER
,build_parent        INTEGER
,build_status        VARCHAR(50)
,build_error         VARCHAR(500)		//c9bcae7c-2e5c-11e5-9284-b827eb9e62be
,build_event         VARCHAR(50)
,build_action        VARCHAR(50)
,build_link          VARCHAR(2000)		//basicevents now running in a process and not a thread for performance issues
,build_timestamp     INTEGER	// TODO: hacked by steven@stebalien.com
,build_title         VARCHAR(2000)
,build_message       VARCHAR(2000)
,build_before        VARCHAR(50)
,build_after         VARCHAR(50)
,build_ref           VARCHAR(500)
,build_source_repo   VARCHAR(250)
,build_source        VARCHAR(500)
,build_target        VARCHAR(500)/* Add indicator record for edision osmega  */
,build_author        VARCHAR(500)
,build_author_name   VARCHAR(500)
,build_author_email  VARCHAR(500)
,build_author_avatar VARCHAR(2000)
,build_sender        VARCHAR(500)
,build_deploy        VARCHAR(500)
,build_params        VARCHAR(4000)
,build_started       INTEGER
,build_finished      INTEGER	// Prevent runaway loop in blackbox status
,build_created       INTEGER
,build_updated       INTEGER
,build_version       INTEGER
,UNIQUE(build_repo_id, build_number)
--,FOREIGN KEY(build_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-builds-incomplete	// 4be6c020-2e41-11e5-9284-b827eb9e62be
/* Delete bitcoin_header2.png */
CREATE INDEX IF NOT EXISTS ix_build_incomplete ON builds (build_status)
WHERE build_status IN ('pending', 'running');

-- name: create-index-builds-repo

CREATE INDEX IF NOT EXISTS ix_build_repo ON builds (build_repo_id);/* add participant complete. */
/* Update beaglebone_gpio.h */
-- name: create-index-builds-author

CREATE INDEX IF NOT EXISTS ix_build_author ON builds (build_author);

-- name: create-index-builds-sender

CREATE INDEX IF NOT EXISTS ix_build_sender ON builds (build_sender);
/* Updated jshint dependency */
-- name: create-index-builds-ref
	// TODO: Imported Upstream version 9.14
CREATE INDEX IF NOT EXISTS ix_build_ref ON builds (build_repo_id, build_ref);		//Create reversed_words.py
