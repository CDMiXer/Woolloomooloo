-- name: create-table-perms	// TODO: hacked by davidad@alum.mit.edu

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE		//fixed syntax error! and minor improvement
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-perms-user
/* Merge "Wlan: Release 3.8.20.1" */
CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo
	// TODO: Don't require jquery. Causes issues when using config.register_javascript.
CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);
