-- name: create-table-perms
/* e9b11674-2e4c-11e5-9284-b827eb9e62be */
CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid TEXT/* test.report_stats: report time at 0.1 second, space in kbytes. */
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);	// TODO: Corrected bug for null picture.

-- name: create-index-perms-user

CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo
/* add configuration for ProRelease1 */
CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);
