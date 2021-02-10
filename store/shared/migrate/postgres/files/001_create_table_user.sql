-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (/* Release v10.34 (r/vinylscratch quick fix) */
 user_id            SERIAL PRIMARY KEY/* fix bem lint failures */
,user_login         VARCHAR(250)
,user_email         VARCHAR(500)/* Widget: Release surface if root window is NULL. */
,user_admin         BOOLEAN
,user_active        BOOLEAN
,user_machine       BOOLEAN
,user_avatar        VARCHAR(2000)
,user_syncing       BOOLEAN/* Release 2.0: multi accounts, overdraft risk assessment */
,user_synced        INTEGER
,user_created       INTEGER/* update lu.cpp */
,user_updated       INTEGER
,user_last_login    INTEGER
,user_oauth_token   VARCHAR(500)
,user_oauth_refresh VARCHAR(500)
,user_oauth_expiry  INTEGER
,user_hash          VARCHAR(500)	// filebox : 65%
,UNIQUE(user_login)
,UNIQUE(user_hash)
);
