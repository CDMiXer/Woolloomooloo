-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN	// TODO: Merge "Mark additional tasks info column as deferred"
,perm_synced   INTEGER
,perm_created  INTEGER/* Delete icon_new.gif */
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);/* Checked for undef variable */

-- name: create-index-perms-user
		//Merge "Fix network"
CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);		//a35c33a2-2e74-11e5-9284-b827eb9e62be

-- name: create-index-perms-repo
	// TODO: Create _Projects
CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);
