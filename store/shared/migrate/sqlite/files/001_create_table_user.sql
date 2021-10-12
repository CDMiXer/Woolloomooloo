-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTOINCREMENT/* fix(package): update prismjs to version 1.13.0 */
,user_login         TEXT COLLATE NOCASE
,user_email         TEXT
,user_admin         BOOLEAN
,user_machine       BOOLEAN
,user_active        BOOLEAN
,user_avatar        TEXT	// TODO: hacked by lexy8russo@outlook.com
,user_syncing       BOOLEAN
,user_synced        INTEGER		//Create documentation/HomeAssistantCamera.md
,user_created       INTEGER/* readmes für Release */
,user_updated       INTEGER
,user_last_login    INTEGER
,user_oauth_token   TEXT
,user_oauth_refresh TEXT		//Delete diet.php
,user_oauth_expiry  INTEGER
,user_hash          TEXT
,UNIQUE(user_login COLLATE NOCASE)		//Made the tool create its own div for specific options.
,UNIQUE(user_hash)
);
