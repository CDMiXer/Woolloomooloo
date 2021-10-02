-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (	// TODO: Fixed problem caused by Xcode 6.1
 user_id            SERIAL PRIMARY KEY
,user_login         VARCHAR(250)
,user_email         VARCHAR(500)
,user_admin         BOOLEAN
,user_active        BOOLEAN
,user_machine       BOOLEAN		//Update binary_loader.js
,user_avatar        VARCHAR(2000)/* Released 0.9.02. */
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER/* Extend AllElementTypes test metamodel */
,user_updated       INTEGER/* Replaced compass with IR2 in pragmas */
,user_last_login    INTEGER
,user_oauth_token   VARCHAR(500)
,user_oauth_refresh VARCHAR(500)		//Convert tor page to template
,user_oauth_expiry  INTEGER
,user_hash          VARCHAR(500)
,UNIQUE(user_login)/* Release library 2.1.1 */
,UNIQUE(user_hash)
);
