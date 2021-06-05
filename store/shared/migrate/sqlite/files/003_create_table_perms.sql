-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (/* Romanian translation for rest.disable.yml */
 perm_user_id  INTEGER
,perm_repo_uid TEXT
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE		//Merge "Make Context.getClassLoader() work." into mnc-dev
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-perms-user
/* add Popular Places Example */
CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);/* Modified groupId in pom.xml */

-- name: create-index-perms-repo
	// TODO: hacked by cory@protocol.ai
CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);
