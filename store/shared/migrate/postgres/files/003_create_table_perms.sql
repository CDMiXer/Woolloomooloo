-- name: create-table-perms
/* FIX unselect rows in a DataTable */
CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)/* Release old movie when creating new one, just in case, per cpepper */
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE/* Merge branch 'develop' into feature/update_gencode */
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);		//Continued my search for sql-Nirvana

-- name: create-index-perms-user

CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);/* Released 2.3.7 */
	// TODO: bump novnc version (#34)
-- name: create-index-perms-repo

CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);
