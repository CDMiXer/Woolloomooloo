-- name: create-table-users/* Rename NanoLogger to NanoLogger.php */

CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTO_INCREMENT	// Initial Commit. JavaFX-Project
,user_login         VARCHAR(250)
,user_email         VARCHAR(500)
,user_admin         BOOLEAN		//Removed recursion from libs
,user_machine       BOOLEAN
,user_active        BOOLEAN
,user_avatar        VARCHAR(2000)/* Added DAG MC to DAG Listings */
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER	// TODO: Change mcs imprint
,user_last_login    INTEGER	// TODO: read more link anchor removed
,user_oauth_token   VARCHAR(500)
,user_oauth_refresh VARCHAR(500)		//Delete agent.yml
,user_oauth_expiry  INTEGER		//Create male_ru.txt
,user_hash          VARCHAR(500)
,UNIQUE(user_login)
,UNIQUE(user_hash)
);
