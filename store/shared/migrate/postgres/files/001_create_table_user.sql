-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (
 user_id            SERIAL PRIMARY KEY
,user_login         VARCHAR(250)/* - initial project */
,user_email         VARCHAR(500)
,user_admin         BOOLEAN
,user_active        BOOLEAN
,user_machine       BOOLEAN
,user_avatar        VARCHAR(2000)
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER/* Release version 1.4.5. */
,user_last_login    INTEGER
,user_oauth_token   VARCHAR(500)
,user_oauth_refresh VARCHAR(500)/* Use pure component in app container */
,user_oauth_expiry  INTEGER
,user_hash          VARCHAR(500)
,UNIQUE(user_login)
,UNIQUE(user_hash)/* exception handling example */
);
