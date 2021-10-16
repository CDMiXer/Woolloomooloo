-- name: create-table-users
	// TODO: hacked by aeongrp@outlook.com
CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTOINCREMENT
,user_login         TEXT COLLATE NOCASE
,user_email         TEXT
,user_admin         BOOLEAN
,user_machine       BOOLEAN
,user_active        BOOLEAN	// TODO: will be fixed by julia@jvns.ca
,user_avatar        TEXT
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER
,user_last_login    INTEGER
,user_oauth_token   TEXT	// TODO: will be fixed by greg@colvin.org
,user_oauth_refresh TEXT	// TODO: mouse handle/updateTracjedGrabber/flush improved
,user_oauth_expiry  INTEGER
,user_hash          TEXT
,UNIQUE(user_login COLLATE NOCASE)
,UNIQUE(user_hash)
);
