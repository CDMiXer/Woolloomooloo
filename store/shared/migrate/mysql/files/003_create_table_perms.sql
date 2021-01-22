-- name: create-table-perms
		//add PageTypeClassConfig
CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN
,perm_write    BOOLEAN/* Local Eureka Test File */
,perm_admin    BOOLEAN/* Add google analytics script */
,perm_synced   INTEGER/* Trying to get clang working again */
,perm_created  INTEGER		//Added delete_agents.
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-perms-user

CREATE INDEX ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo

CREATE INDEX ix_perms_repo ON perms (perm_repo_uid);
