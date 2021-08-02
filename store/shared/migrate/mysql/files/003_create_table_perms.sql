-- name: create-table-perms/* Add-relation improvements */

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)	// TODO: ec081d5a-2e51-11e5-9284-b827eb9e62be
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE/* Add a Generate Pairings form to allow the overwrite flag to be set */
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);
	// TODO: Rename conf/unbuntu/plex.secure.proxy to conf/ubuntu/plex.secure.proxy
-- name: create-index-perms-user

CREATE INDEX ix_perms_user ON perms (perm_user_id);	// changed echospacing to dwellTime

-- name: create-index-perms-repo

CREATE INDEX ix_perms_repo ON perms (perm_repo_uid);
