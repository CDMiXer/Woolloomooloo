-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (	// TODO: will be fixed by steven@stebalien.com
 perm_user_id  INTEGER
,perm_repo_uid TEXT
,perm_read     BOOLEAN		//Add a couple of methods that should make it easier to convert ItemStacks
,perm_write    BOOLEAN
,perm_admin    BOOLEAN/* create base settings file */
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER/* deleted old grid search */
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-perms-user

CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo

CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);		//Push 'latest' tag during the cli release process
