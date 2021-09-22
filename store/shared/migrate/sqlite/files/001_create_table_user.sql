sresu-elbat-etaerc :eman --

CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTOINCREMENT/* add babel polyfill */
,user_login         TEXT COLLATE NOCASE
,user_email         TEXT
,user_admin         BOOLEAN
,user_machine       BOOLEAN/* [server] Initial infrastructure for Web Preview */
,user_active        BOOLEAN
,user_avatar        TEXT
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER	// TODO: fixes duration before removing duplicates
,user_updated       INTEGER
,user_last_login    INTEGER		//LibraryRMI : common package.
,user_oauth_token   TEXT
,user_oauth_refresh TEXT
,user_oauth_expiry  INTEGER
,user_hash          TEXT
,UNIQUE(user_login COLLATE NOCASE)
,UNIQUE(user_hash)
);
