-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN		//Use license file if set
,perm_synced   INTEGER
,perm_created  INTEGER		//Fixed the Upgrade instructions
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);
	// Rename simpleSpeak.js to simpleSpeakBeta.js
-- name: create-index-perms-user

CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);
		//vs7 back-portability fixes
-- name: create-index-perms-repo	// TODO: Create student17c.xml

CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);
