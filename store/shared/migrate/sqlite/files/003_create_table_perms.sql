-- name: create-table-perms	// TODO: [test] added extension tests for user update

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid TEXT		//bind parameter might be removed in future
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER/* Released version 1.0.0. */
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-perms-user	// TODO: will be fixed by sjors@sprovoost.nl

CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo/* Automatic changelog generation #4413 [ci skip] */

CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);
