-- name: create-table-perms
/* [artifactory-release] Release version 0.8.8.RELEASE */
CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid TEXT	// TODO: hacked by 13860583249@yeah.net
,perm_read     BOOLEAN
,perm_write    BOOLEAN	// TODO: hacked by zaq1tomo@gmail.com
,perm_admin    BOOLEAN/* Release DBFlute-1.1.0-RC2 */
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER		//javaee7 archetype prototypes
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-perms-user

CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo/* Release for 19.0.0 */

CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);
