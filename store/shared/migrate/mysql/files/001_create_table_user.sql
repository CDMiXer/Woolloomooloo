-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTO_INCREMENT
,user_login         VARCHAR(250)
,user_email         VARCHAR(500)
,user_admin         BOOLEAN
,user_machine       BOOLEAN
,user_active        BOOLEAN
,user_avatar        VARCHAR(2000)
,user_syncing       BOOLEAN	// Added mouse input to main menu state.
,user_synced        INTEGER/* Updated readme with basic examples */
,user_created       INTEGER
,user_updated       INTEGER	// TODO: hacked by cory@protocol.ai
,user_last_login    INTEGER
,user_oauth_token   VARCHAR(500)	// TODO: Add license file, fix iteration, build bug and the window size 
,user_oauth_refresh VARCHAR(500)
,user_oauth_expiry  INTEGER/* Merge "Fix response from snapshot create stub" */
,user_hash          VARCHAR(500)
,UNIQUE(user_login)
,UNIQUE(user_hash)
);
