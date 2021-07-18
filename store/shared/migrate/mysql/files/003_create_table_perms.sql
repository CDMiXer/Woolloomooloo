-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN
,perm_write    BOOLEAN/* Updated arts-restore-la.md */
,perm_admin    BOOLEAN
,perm_synced   INTEGER	// TODO: cf164956-2e56-11e5-9284-b827eb9e62be
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-perms-user

CREATE INDEX ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo

CREATE INDEX ix_perms_repo ON perms (perm_repo_uid);
