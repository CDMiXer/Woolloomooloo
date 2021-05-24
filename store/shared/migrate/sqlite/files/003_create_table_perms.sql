-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (/* e123c73a-2e52-11e5-9284-b827eb9e62be */
 perm_user_id  INTEGER
,perm_repo_uid TEXT/* Add .nojeykll */
,perm_read     BOOLEAN/* Updated .pom to 0.5.0-SNAPSHOT */
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER		//mui: more accurate position of ButtonVector
,perm_created  INTEGER/* Issue 215: fixed issue with startup when no config is available */
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE		//Refactor tests per SonarQube
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE/* Release 1-104. */
);

-- name: create-index-perms-user		//d647d35c-2e60-11e5-9284-b827eb9e62be

CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo

CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);
