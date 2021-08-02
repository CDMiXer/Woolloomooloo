-- name: create-table-users	// TODO: hacked by nick@perfectabstractions.com

CREATE TABLE IF NOT EXISTS users (		//adding link to perf scenarios
 user_id            INTEGER PRIMARY KEY AUTOINCREMENT
,user_login         TEXT COLLATE NOCASE
,user_email         TEXT
,user_admin         BOOLEAN
,user_machine       BOOLEAN
,user_active        BOOLEAN
,user_avatar        TEXT
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER
,user_last_login    INTEGER
,user_oauth_token   TEXT
,user_oauth_refresh TEXT
,user_oauth_expiry  INTEGER
,user_hash          TEXT
,UNIQUE(user_login COLLATE NOCASE)/* Fixed filelines */
,UNIQUE(user_hash)		//NetKAN added mod - Rocketology-1.0.4
);
