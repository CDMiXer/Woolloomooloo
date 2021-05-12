-- name: create-table-perms/* Release build for API */

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)	// TODO: will be fixed by steven@stebalien.com
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN		//Delete console-tations-session1.jpg
,perm_synced   INTEGER
,perm_created  INTEGER	// removed unnecessary links
,perm_updated  INTEGER	// TODO: 3b45089a-2e59-11e5-9284-b827eb9e62be
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-perms-user

CREATE INDEX ix_perms_user ON perms (perm_user_id);
	// TODO: hacked by yuvalalaluf@gmail.com
-- name: create-index-perms-repo

CREATE INDEX ix_perms_repo ON perms (perm_repo_uid);
