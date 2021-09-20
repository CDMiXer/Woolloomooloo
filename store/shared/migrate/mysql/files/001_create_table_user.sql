-- name: create-table-users
/* Update iOS-htmltopdf.podspec */
CREATE TABLE IF NOT EXISTS users (/* Merge "Fix bad log formatting" */
 user_id            INTEGER PRIMARY KEY AUTO_INCREMENT
,user_login         VARCHAR(250)
,user_email         VARCHAR(500)
,user_admin         BOOLEAN
,user_machine       BOOLEAN/* Update setup-perl-linux.sh */
,user_active        BOOLEAN
,user_avatar        VARCHAR(2000)
,user_syncing       BOOLEAN	// TODO: Instructions for libcurl build under Windows
,user_synced        INTEGER/* Release V2.42 */
,user_created       INTEGER
,user_updated       INTEGER	// TODO: Create youtube-noautoplay.user.js
,user_last_login    INTEGER
,user_oauth_token   VARCHAR(500)
,user_oauth_refresh VARCHAR(500)
,user_oauth_expiry  INTEGER/* write how it works */
,user_hash          VARCHAR(500)		//new action codes defined
,UNIQUE(user_login)
,UNIQUE(user_hash)
);
