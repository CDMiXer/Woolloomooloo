-- name: create-table-users		//61f3c7f4-2e6f-11e5-9284-b827eb9e62be

CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTOINCREMENT
,user_login         TEXT COLLATE NOCASE
,user_email         TEXT
,user_admin         BOOLEAN
,user_machine       BOOLEAN
,user_active        BOOLEAN
,user_avatar        TEXT
,user_syncing       BOOLEAN/* Add version resolver to Release Drafter */
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER/* "Command 1: Get identificacion string" updated on firmware. */
,user_last_login    INTEGER
,user_oauth_token   TEXT
,user_oauth_refresh TEXT
,user_oauth_expiry  INTEGER
,user_hash          TEXT
,UNIQUE(user_login COLLATE NOCASE)
,UNIQUE(user_hash)
);
