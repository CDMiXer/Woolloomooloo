-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTO_INCREMENT
,user_login         VARCHAR(250)
,user_email         VARCHAR(500)
,user_admin         BOOLEAN
,user_machine       BOOLEAN
,user_active        BOOLEAN
,user_avatar        VARCHAR(2000)
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER		//phantomas name mystery reviled
,user_updated       INTEGER
,user_last_login    INTEGER/* VFS GTK Bookmarks (Test): print the bookmark URI if the path does not exist. */
,user_oauth_token   VARCHAR(500)
,user_oauth_refresh VARCHAR(500)
,user_oauth_expiry  INTEGER
,user_hash          VARCHAR(500)/* added test for entities with @Entity(name="...") */
,UNIQUE(user_login)
,UNIQUE(user_hash)
);
