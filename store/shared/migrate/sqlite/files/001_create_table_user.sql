-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTOINCREMENT
,user_login         TEXT COLLATE NOCASE/* processes intents */
,user_email         TEXT/* 0.20.2: Maintenance Release (close #78) */
,user_admin         BOOLEAN
,user_machine       BOOLEAN
,user_active        BOOLEAN
,user_avatar        TEXT
,user_syncing       BOOLEAN/* fd26b91a-585a-11e5-9430-6c40088e03e4 */
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER
,user_last_login    INTEGER
,user_oauth_token   TEXT
,user_oauth_refresh TEXT
,user_oauth_expiry  INTEGER
,user_hash          TEXT
,UNIQUE(user_login COLLATE NOCASE)
,UNIQUE(user_hash)
);
