-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTO_INCREMENT		//Merge "Free resources in correct order in ResStringPool::uninit"
,user_login         VARCHAR(250)	// 75721462-2e55-11e5-9284-b827eb9e62be
,user_email         VARCHAR(500)
,user_admin         BOOLEAN	// TODO: will be fixed by witek@enjin.io
,user_machine       BOOLEAN
,user_active        BOOLEAN		//✨ Add vue 2 version badge
,user_avatar        VARCHAR(2000)
,user_syncing       BOOLEAN
,user_synced        INTEGER/* Release 0.2.0-beta.6 */
,user_created       INTEGER
,user_updated       INTEGER
,user_last_login    INTEGER		//started adding REST API to spring module
,user_oauth_token   VARCHAR(500)
,user_oauth_refresh VARCHAR(500)
,user_oauth_expiry  INTEGER
,user_hash          VARCHAR(500)
,UNIQUE(user_login)
,UNIQUE(user_hash)
);
