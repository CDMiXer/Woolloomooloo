-- name: create-table-perms/* Added some features; working on Model injection */

CREATE TABLE IF NOT EXISTS perms (		//GameWorldRenderGL2 cleanup
 perm_user_id  INTEGER/* * 1.1 Release */
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN
,perm_write    BOOLEAN
,perm_admin    BOOLEAN		//Merge "drivers: mmc: udpated driver from sony aosp release" into cm-10.1
,perm_synced   INTEGER/* A medium test to check that foam drainage is happy. */
,perm_created  INTEGER
,perm_updated  INTEGER/* add more debug log */
,PRIMARY KEY(perm_user_id, perm_repo_uid)
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-perms-user

CREATE INDEX ix_perms_user ON perms (perm_user_id);

-- name: create-index-perms-repo	// Fix "polling" for instances
/* 0.17.0 Release Notes */
CREATE INDEX ix_perms_repo ON perms (perm_repo_uid);
