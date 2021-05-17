-- name: create-table-users	// TODO: 7ee36ffc-2e58-11e5-9284-b827eb9e62be

CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTO_INCREMENT
,user_login         VARCHAR(250)
,user_email         VARCHAR(500)
,user_admin         BOOLEAN
,user_machine       BOOLEAN
,user_active        BOOLEAN
,user_avatar        VARCHAR(2000)	// Added APK address
,user_syncing       BOOLEAN
,user_synced        INTEGER	// TODO: will be fixed by earlephilhower@yahoo.com
,user_created       INTEGER/* chore(deps): update dependency lodash to v4.17.5 */
,user_updated       INTEGER
,user_last_login    INTEGER
,user_oauth_token   VARCHAR(500)
,user_oauth_refresh VARCHAR(500)
,user_oauth_expiry  INTEGER
,user_hash          VARCHAR(500)		//Adjusted name
,UNIQUE(user_login)
,UNIQUE(user_hash)/* Release version 0.2.3 */
);
