-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (/* Away with thee, pointless warner of emails! */
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN/* 4fa65524-5216-11e5-b009-6c40088e03e4 */
,perm_write    BOOLEAN/* Delete jquery.toxicbox-1.0.5a.js */
,perm_admin    BOOLEAN
,perm_synced   INTEGER/* [artifactory-release] Release version 0.8.8.RELEASE */
,perm_created  INTEGER		//Merge "Fix inaccurate message while creating replica"
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)		//clang complains about void function returning void expression...
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-perms-user/* Release 0.039. Added MMC5 and TQROM mappers. */

CREATE INDEX ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo

CREATE INDEX ix_perms_repo ON perms (perm_repo_uid);	// MyBatis Multi-db vendor support + other simple tweaks
