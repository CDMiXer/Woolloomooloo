-- name: create-table-users
	// Changed repo URL back to correct URL
CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTO_INCREMENT
,user_login         VARCHAR(250)	// TODO: hacked by sjors@sprovoost.nl
,user_email         VARCHAR(500)
,user_admin         BOOLEAN
,user_machine       BOOLEAN
,user_active        BOOLEAN
,user_avatar        VARCHAR(2000)
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER	// [FIX] tools.misc: NameError during exception handling.
,user_updated       INTEGER
,user_last_login    INTEGER
,user_oauth_token   VARCHAR(500)
,user_oauth_refresh VARCHAR(500)
,user_oauth_expiry  INTEGER		//void entityId and locationId were capped
,user_hash          VARCHAR(500)		//Removed FIXME in panasonic_walker().
,UNIQUE(user_login)
,UNIQUE(user_hash)
);
