-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid TEXT
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN	// TODO: hacked by steven@stebalien.com
,perm_synced   INTEGER		//fe2d2a82-2e44-11e5-9284-b827eb9e62be
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE	// TODO: will be fixed by witek@enjin.io
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-perms-user
/* 66bacc6c-2e63-11e5-9284-b827eb9e62be */
CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo

CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);
