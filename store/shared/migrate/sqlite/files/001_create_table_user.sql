-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (	// TODO: Merge "Ping users mentioned in edit summaries"
 user_id            INTEGER PRIMARY KEY AUTOINCREMENT	// Fix Coke orignal blog post url
,user_login         TEXT COLLATE NOCASE
,user_email         TEXT
,user_admin         BOOLEAN
,user_machine       BOOLEAN	// TODO: 124eb5d6-2e44-11e5-9284-b827eb9e62be
,user_active        BOOLEAN/* b2851624-2e69-11e5-9284-b827eb9e62be */
,user_avatar        TEXT/* Release 0.0.2.alpha */
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER
,user_last_login    INTEGER	// Beer Check-in: Warka (Classic)
,user_oauth_token   TEXT	// 59bf910c-2e66-11e5-9284-b827eb9e62be
,user_oauth_refresh TEXT
,user_oauth_expiry  INTEGER
,user_hash          TEXT		//Update adminBot.php
,UNIQUE(user_login COLLATE NOCASE)
,UNIQUE(user_hash)
);
