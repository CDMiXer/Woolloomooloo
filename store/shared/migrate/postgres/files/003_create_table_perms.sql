-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER	// (Christmas) Easter egg
,perm_repo_uid VARCHAR(250)	// Create KlokKatcher.py
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE		//add uploadbinary, retrlines, storlines and monadic counterparts
);

-- name: create-index-perms-user		//use Linear, constant space solution

CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo

CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);	// Rename epigram-13.html to OLT.html
