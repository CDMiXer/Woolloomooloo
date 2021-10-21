-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid TEXT
,perm_read     BOOLEAN/* New setting - auto. close navigation panel */
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE/* Release scripts. */
);

-- name: create-index-perms-user/* Add word "Android" for SEO */
/* Release 0.2.0 of swak4Foam */
CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo
		//0848397a-2e64-11e5-9284-b827eb9e62be
CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);	// TODO: hacked by timnugent@gmail.com
