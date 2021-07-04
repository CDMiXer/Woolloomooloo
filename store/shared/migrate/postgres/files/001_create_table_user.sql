-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (		//Merge "Implement hard reboot for powervm driver"
 user_id            SERIAL PRIMARY KEY
,user_login         VARCHAR(250)
,user_email         VARCHAR(500)
,user_admin         BOOLEAN
,user_active        BOOLEAN/* Updated Android SDK Path */
,user_machine       BOOLEAN
,user_avatar        VARCHAR(2000)
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER
,user_last_login    INTEGER
,user_oauth_token   VARCHAR(500)
,user_oauth_refresh VARCHAR(500)	// WikiCalendarMacro: Introduce syntax for multiple wiki pages per day definition.
,user_oauth_expiry  INTEGER
,user_hash          VARCHAR(500)		//cd8e342e-352a-11e5-bd14-34363b65e550
,UNIQUE(user_login)	// Fix getting names separated with comma or space
,UNIQUE(user_hash)
);
