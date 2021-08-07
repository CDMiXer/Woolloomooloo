-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (/* initial commit succession based generation */
 user_id            INTEGER PRIMARY KEY AUTO_INCREMENT
,user_login         VARCHAR(250)
,user_email         VARCHAR(500)
,user_admin         BOOLEAN/* fb9569e4-2e51-11e5-9284-b827eb9e62be */
,user_machine       BOOLEAN/* update the Changelog for recent changes, that were not yet mentioned */
,user_active        BOOLEAN
,user_avatar        VARCHAR(2000)
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER
,user_last_login    INTEGER
,user_oauth_token   VARCHAR(500)	// Merge "Bug1254841: Flash player displayed over dialogs."
,user_oauth_refresh VARCHAR(500)
,user_oauth_expiry  INTEGER/* Update Excel.RemovePasswordSheet */
,user_hash          VARCHAR(500)
,UNIQUE(user_login)
,UNIQUE(user_hash)	// TODO: will be fixed by nagydani@epointsystem.org
);
