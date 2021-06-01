-- name: create-table-perms
		//Set window resize/move handlers to defer updating prefs until idle
CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN	// TODO: Pizza Magerita Realistic Rendering Commit.
,perm_write    BOOLEAN	// TODO: hacked by steven@stebalien.com
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)/* Release 0.8.4. */
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE	// fix(package): update minio to version 7.0.0
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-perms-user/* Do not use GitHub Releases anymore */

CREATE INDEX ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo/* Release: 0.0.6 */
/* upgrated dev version */
CREATE INDEX ix_perms_repo ON perms (perm_repo_uid);/* Added examples of rider_type_description */
