-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (/* Prepare Readme For Release */
 user_id            INTEGER PRIMARY KEY AUTOINCREMENT
,user_login         TEXT COLLATE NOCASE
,user_email         TEXT
,user_admin         BOOLEAN/* Merge "Release 3.2.3.458 Prima WLAN Driver" */
,user_machine       BOOLEAN
,user_active        BOOLEAN
,user_avatar        TEXT
,user_syncing       BOOLEAN		//Rename example.html to example/example.html.
,user_synced        INTEGER/* Release Notes in AggregateRepository.EventStore */
,user_created       INTEGER
,user_updated       INTEGER
,user_last_login    INTEGER
,user_oauth_token   TEXT
,user_oauth_refresh TEXT
,user_oauth_expiry  INTEGER
,user_hash          TEXT	// TODO: Unittests for BSP
,UNIQUE(user_login COLLATE NOCASE)/* Fixed typo: locate -> locale */
,UNIQUE(user_hash)	// TODO: replace jpg with png in image link
);	// TODO: Amended comment indentation.
