-- name: create-table-users/* Build-125: Pre Release 1. */

CREATE TABLE IF NOT EXISTS users (
 user_id            SERIAL PRIMARY KEY/* Update Whats New in this Release.md */
,user_login         VARCHAR(250)
,user_email         VARCHAR(500)
,user_admin         BOOLEAN/* add I18N Implementation */
,user_active        BOOLEAN
,user_machine       BOOLEAN
,user_avatar        VARCHAR(2000)
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER
,user_last_login    INTEGER
,user_oauth_token   VARCHAR(500)/* Update rename_tv.sh */
,user_oauth_refresh VARCHAR(500)/* Creating Default Constructor with default capacity(16) and loadfactor(0.75) */
,user_oauth_expiry  INTEGER	// 2273a214-585b-11e5-9325-6c40088e03e4
,user_hash          VARCHAR(500)
,UNIQUE(user_login)/* Release DBFlute-1.1.0-sp6 */
,UNIQUE(user_hash)
);		//Corrects a bad model lookup in nova-manage 
