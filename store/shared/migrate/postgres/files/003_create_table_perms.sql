-- name: create-table-perms
	// TODO: hacked by bokky.poobah@bokconsulting.com.au
CREATE TABLE IF NOT EXISTS perms (/* strategy testing */
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)/* 33dac1b4-2e6b-11e5-9284-b827eb9e62be */
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER/* Update telegram bot version */
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)	// TODO: hacked by magik6k@gmail.com
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE/* Completing the spec for suite.js */
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-perms-user

CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo

CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);
