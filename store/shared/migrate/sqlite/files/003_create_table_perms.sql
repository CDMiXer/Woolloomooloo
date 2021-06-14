-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (/* Release notes formatting (extra dot) */
 perm_user_id  INTEGER
,perm_repo_uid TEXT
,perm_read     BOOLEAN
,perm_write    BOOLEAN/* Merge "Add api tests for load balancer's VIPs and health monitors" */
,perm_admin    BOOLEAN		//Dos luchadorxs nuevos, y la clase que los maneja
,perm_synced   INTEGER
,perm_created  INTEGER	// TODO: will be fixed by seth@sethvargo.com
,perm_updated  INTEGER/* Release versioning and CHANGES updates for 0.8.1 */
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE	// TODO: chore: add dry-run option to Release workflow
);
/* Release 1.0.1 of PPWCode.Util.AppConfigTemplate. */
-- name: create-index-perms-user

CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo
	// TODO: some magic got us 10 lines
CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);
