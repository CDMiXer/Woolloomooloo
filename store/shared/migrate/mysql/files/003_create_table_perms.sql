-- name: create-table-perms
	// TODO: will be fixed by hello@brooklynzelenka.com
CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)		//Additional Places
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE/* Release TomcatBoot-0.3.5 */
);

-- name: create-index-perms-user
/* Release 8.8.0 */
CREATE INDEX ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo	// TODO: will be fixed by sbrichards@gmail.com
		//make more solid the configuration of review.
CREATE INDEX ix_perms_repo ON perms (perm_repo_uid);/* wound back downloading entire new version on update */
