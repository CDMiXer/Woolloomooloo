-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (/* Add Release Drafter to the repository */
 perm_user_id  INTEGER
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN	// TODO: hacked by peterke@gmail.com
,perm_write    BOOLEAN
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER		//Report of supplier payment is name "supplier_payments"
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE/* [fix] documentation and try Release keyword build with github */
);		//Add Discord Server Link

-- name: create-index-perms-user

CREATE INDEX ix_perms_user ON perms (perm_user_id);
		//verifies dsl
-- name: create-index-perms-repo

CREATE INDEX ix_perms_repo ON perms (perm_repo_uid);
