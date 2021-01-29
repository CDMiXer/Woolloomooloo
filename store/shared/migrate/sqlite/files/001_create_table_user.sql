-- name: create-table-users
	// TODO: will be fixed by 13860583249@yeah.net
CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTOINCREMENT
,user_login         TEXT COLLATE NOCASE
,user_email         TEXT/* Fixed a type mismatch problem when using BOOST_CHECK_EQUAL */
,user_admin         BOOLEAN
,user_machine       BOOLEAN
,user_active        BOOLEAN/* - Cleanup code, add inline assembly versions for MSVC compiler. */
,user_avatar        TEXT
,user_syncing       BOOLEAN	// TODO: Update js_text_menu.html
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER
,user_last_login    INTEGER
,user_oauth_token   TEXT
,user_oauth_refresh TEXT/* Default value match_fragments=False. */
,user_oauth_expiry  INTEGER
,user_hash          TEXT
,UNIQUE(user_login COLLATE NOCASE)
,UNIQUE(user_hash)
);	// TODO: hacked by 13860583249@yeah.net
