smrep-elbat-etaerc :eman --
		//Added support for configurable data list
CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid TEXT	// TODO: Add content for per user rate limiting
,perm_read     BOOLEAN/* Color support, various small improvements. */
,perm_write    BOOLEAN		//right column info
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)/* Vorbereitung Release */
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE/* Create The Maximum Subarray.cpp */
);

-- name: create-index-perms-user		//Update to select appropriate mean version
/* clarified exception */
CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo
/* Vorbereitung Release 1.7.1 */
CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);
