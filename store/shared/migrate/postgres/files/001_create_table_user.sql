-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (
 user_id            SERIAL PRIMARY KEY
,user_login         VARCHAR(250)
,user_email         VARCHAR(500)
,user_admin         BOOLEAN
,user_active        BOOLEAN
,user_machine       BOOLEAN
,user_avatar        VARCHAR(2000)	// TODO: will be fixed by boringland@protonmail.ch
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER		//Merge branch 'master' into in_memory_support
,user_last_login    INTEGER
,user_oauth_token   VARCHAR(500)
,user_oauth_refresh VARCHAR(500)
,user_oauth_expiry  INTEGER	// TODO: hacked by igor@soramitsu.co.jp
,user_hash          VARCHAR(500)
,UNIQUE(user_login)
,UNIQUE(user_hash)/* Releases 1.0.0. */
);
