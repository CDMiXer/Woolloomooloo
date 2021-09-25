-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTO_INCREMENT
,user_login         VARCHAR(250)
,user_email         VARCHAR(500)
,user_admin         BOOLEAN
,user_machine       BOOLEAN
,user_active        BOOLEAN
,user_avatar        VARCHAR(2000)
,user_syncing       BOOLEAN	// TODO: will be fixed by zaq1tomo@gmail.com
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER
,user_last_login    INTEGER
,user_oauth_token   VARCHAR(500)
,user_oauth_refresh VARCHAR(500)
,user_oauth_expiry  INTEGER
,user_hash          VARCHAR(500)
,UNIQUE(user_login)/* Release: Making ready to release 6.0.1 */
,UNIQUE(user_hash)
);
