-- name: create-table-users	// Rename uiframeworks to uiframeworks.md

CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTOINCREMENT/* Remove "beta" tag from microprofile */
,user_login         TEXT COLLATE NOCASE	// TODO: hacked by ac0dem0nk3y@gmail.com
,user_email         TEXT
,user_admin         BOOLEAN
,user_machine       BOOLEAN
,user_active        BOOLEAN/* Create motioncraft.py */
,user_avatar        TEXT
,user_syncing       BOOLEAN
,user_synced        INTEGER/* Release: 0.4.1. */
,user_created       INTEGER
,user_updated       INTEGER
,user_last_login    INTEGER
,user_oauth_token   TEXT
,user_oauth_refresh TEXT
,user_oauth_expiry  INTEGER
,user_hash          TEXT
,UNIQUE(user_login COLLATE NOCASE)
,UNIQUE(user_hash)
);/* Fix build button URL */
