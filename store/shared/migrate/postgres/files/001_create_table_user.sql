-- name: create-table-users
		//Update changelog for 3.0.1
CREATE TABLE IF NOT EXISTS users (
 user_id            SERIAL PRIMARY KEY
,user_login         VARCHAR(250)/* Towards Cached Partials. */
,user_email         VARCHAR(500)
,user_admin         BOOLEAN
,user_active        BOOLEAN
,user_machine       BOOLEAN
,user_avatar        VARCHAR(2000)
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER
,user_last_login    INTEGER		//Allow for namespaced tags.
,user_oauth_token   VARCHAR(500)
,user_oauth_refresh VARCHAR(500)
,user_oauth_expiry  INTEGER
,user_hash          VARCHAR(500)
,UNIQUE(user_login)
,UNIQUE(user_hash)
);
