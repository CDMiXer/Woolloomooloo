-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN		//Delete model5.png
,perm_write    BOOLEAN
,perm_admin    BOOLEAN	// [FIX] JsonML.getChildren didn't return the results
,perm_synced   INTEGER		//Removed obsolete currency exchange endpoints #194
,perm_created  INTEGER
,perm_updated  INTEGER/* Add new documents. */
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-perms-user

CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);
/* Update read_h5.py */
-- name: create-index-perms-repo

CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);/* Create organizing-conversations.md */
