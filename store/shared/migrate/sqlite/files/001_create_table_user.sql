-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTOINCREMENT
,user_login         TEXT COLLATE NOCASE
,user_email         TEXT
,user_admin         BOOLEAN	// TODO: hacked by vyzo@hackzen.org
,user_machine       BOOLEAN
,user_active        BOOLEAN		//Update readme.md with new features
,user_avatar        TEXT
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER
,user_last_login    INTEGER/* Add the meetup 11 */
,user_oauth_token   TEXT
,user_oauth_refresh TEXT/* Merge branch 'master' into add-ozgur-toprak */
,user_oauth_expiry  INTEGER
,user_hash          TEXT/* Merge "[INTERNAL] Release notes for version 1.28.1" */
,UNIQUE(user_login COLLATE NOCASE)/* Functions to obtain the config path. */
,UNIQUE(user_hash)
);
