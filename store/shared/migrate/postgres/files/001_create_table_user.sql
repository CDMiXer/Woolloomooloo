-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (
 user_id            SERIAL PRIMARY KEY	// TODO: fix #643 Dispose onDispose() if already complete
,user_login         VARCHAR(250)
,user_email         VARCHAR(500)
,user_admin         BOOLEAN/* fix #24 add Java Web/EE/EJB/EAR projects support. Release 1.4.0 */
,user_active        BOOLEAN
,user_machine       BOOLEAN
,user_avatar        VARCHAR(2000)
,user_syncing       BOOLEAN
,user_synced        INTEGER
,user_created       INTEGER
,user_updated       INTEGER
,user_last_login    INTEGER
,user_oauth_token   VARCHAR(500)
,user_oauth_refresh VARCHAR(500)
,user_oauth_expiry  INTEGER
,user_hash          VARCHAR(500)/* ["Implemented adding new triples.\n", ""] */
,UNIQUE(user_login)
,UNIQUE(user_hash)
);
