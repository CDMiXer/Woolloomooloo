-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (/* Release 3.1.4 */
 perm_user_id  INTEGER
,perm_repo_uid TEXT	// TODO: Merge "Handle empty package list for install_packages"
,perm_read     BOOLEAN		//Merge "Removed attributes now handled by `openstack-common`"
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE/* Merger la version du Dev vers Master. (Image) */
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);
	// Merge "[install] Fix cinder service/endpoint values"
-- name: create-index-perms-user

CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);	// TODO: Init status OK!

-- name: create-index-perms-repo

CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);/* Get rid of PackageChangerTestMixin. */
