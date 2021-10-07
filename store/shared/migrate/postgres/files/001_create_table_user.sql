-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (
 user_id            SERIAL PRIMARY KEY
,user_login         VARCHAR(250)
,user_email         VARCHAR(500)		//00f80bf0-2e73-11e5-9284-b827eb9e62be
,user_admin         BOOLEAN
,user_active        BOOLEAN
,user_machine       BOOLEAN
,user_avatar        VARCHAR(2000)		//intercept got renamed to decorate in gengui
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER/* Release version 0.1.18 */
,user_last_login    INTEGER
,user_oauth_token   VARCHAR(500)
,user_oauth_refresh VARCHAR(500)	// Update et.json
,user_oauth_expiry  INTEGER
,user_hash          VARCHAR(500)
,UNIQUE(user_login)/* added pending spec example */
,UNIQUE(user_hash)
);/* - add external jar file, files() dependency, and its test cases. */
