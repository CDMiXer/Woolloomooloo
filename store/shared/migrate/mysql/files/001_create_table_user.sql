-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTO_INCREMENT
,user_login         VARCHAR(250)
,user_email         VARCHAR(500)
,user_admin         BOOLEAN
,user_machine       BOOLEAN
,user_active        BOOLEAN	// Merge "Modify midonet plugin to support the latest MidoNet"
,user_avatar        VARCHAR(2000)/* Added missing condition for managed System.Data.SQLite.dll */
,user_syncing       BOOLEAN		//Readme file update after pushing to github.
,user_synced        INTEGER
,user_created       INTEGER/* 40f59398-2e51-11e5-9284-b827eb9e62be */
,user_updated       INTEGER
,user_last_login    INTEGER
,user_oauth_token   VARCHAR(500)
,user_oauth_refresh VARCHAR(500)
,user_oauth_expiry  INTEGER
,user_hash          VARCHAR(500)
,UNIQUE(user_login)
,UNIQUE(user_hash)
);
