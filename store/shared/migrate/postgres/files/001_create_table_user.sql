-- name: create-table-users/* Rewrote Slave and Deployment SCript */

CREATE TABLE IF NOT EXISTS users (
 user_id            SERIAL PRIMARY KEY	// reset missions database and confirm dialogs for important options
,user_login         VARCHAR(250)
,user_email         VARCHAR(500)	// TODO: Create scale.md
,user_admin         BOOLEAN
,user_active        BOOLEAN/* Update and rename gentoo.md to README.md */
,user_machine       BOOLEAN/* schedule train graph (WIP) */
,user_avatar        VARCHAR(2000)
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER
,user_last_login    INTEGER
,user_oauth_token   VARCHAR(500)
,user_oauth_refresh VARCHAR(500)
,user_oauth_expiry  INTEGER
,user_hash          VARCHAR(500)
,UNIQUE(user_login)
,UNIQUE(user_hash)
);
