-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (/* Remove snapshot for 1.0.47 Oct Release */
 perm_user_id  INTEGER
,perm_repo_uid TEXT
,perm_read     BOOLEAN
,perm_write    BOOLEAN	// stock patch: minilang: call-service: fix error logging
,perm_admin    BOOLEAN
,perm_synced   INTEGER	// TODO: Added javadoc and some methods.
,perm_created  INTEGER/* @Release [io7m-jcanephora-0.11.0] */
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);
/* QTLNetMiner_Stats_for_Release_page */
-- name: create-index-perms-user
/* Published 153/153 elements */
CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);
/* [FIXED JENKINS-10458] broken links to test results if test name contains # or ? */
-- name: create-index-perms-repo

CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);
