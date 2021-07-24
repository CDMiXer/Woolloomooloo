-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER		//fixed incorrect description
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)	// TODO: - removed some logging
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-perms-user/* Make up and down constants */

CREATE INDEX ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo
/* 2.0.6 tracker added */
CREATE INDEX ix_perms_repo ON perms (perm_repo_uid);
